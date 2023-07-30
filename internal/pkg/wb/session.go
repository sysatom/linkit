package wb

import (
	"encoding/json"
	"github.com/gorilla/websocket"
	"github.com/sysatom/linkit/internal/pkg/logs"
	"github.com/sysatom/linkit/internal/pkg/types"
	"sync"
	"sync/atomic"
	"time"
)

const (
	// idleSessionTimeout defines duration of being idle before terminating a session.
	idleSessionTimeout = time.Second * 55

	// defaultMaxMessageSize is the default maximum message size
	defaultMaxMessageSize = 1 << 19 // 512K

	// Time allowed to write a message to the peer.
	writeWait = 10 * time.Second

	// Time allowed to read the next pong message from the peer.
	pongWait = idleSessionTimeout

	// Send pings to peer with this period. Must be less than pongWait.
	pingPeriod = (pongWait * 9) / 10

	// Maximum number of queued messages before session is considered stale and dropped.
	sendQueueLimit = 128
)

// ProxyReqType is the type of proxy requests.
type ProxyReqType int

// Session represents a single WS connection or a long polling session. A user may have multiple
// sessions.
type Session struct {
	// Session ID
	sid string

	// Websocket. Set only for websocket sessions.
	ws *websocket.Conn

	// Protocol version of the client: ((major & 0xff) << 8) | (minor & 0xff).
	ver int

	// Device ID of the client
	deviceID string
	// Platform: web, ios, android
	platform string
	// Human language of the client
	lang string
	// Country code of the client
	countryCode string

	// Time when the long polling session was last refreshed
	lastTouched time.Time

	// Time when the session received any packer from client
	lastAction int64

	// Timer which triggers after some seconds to mark background session as foreground.
	bkgTimer *time.Timer

	// Synchronizes access to session store in cluster mode:
	// subscribe/unsubscribe replies are asynchronous.
	sessionStoreLock sync.Mutex
	// Indicates that the session is terminating.
	// After this flag's been flipped to true, there must not be anymore writes
	// into the session's send channel.
	// Read/written atomically.
	// 0 = false
	// 1 = true
	terminating int32

	// Background session: subscription presence notifications and online status are delayed.
	background bool

	// Outbound messages, buffered.
	// The content must be serialized in format suitable for the session.
	send chan any

	// Channel for shutting down the session, buffer 1.
	// Content in the same format as for 'send'
	stop chan any

	// detach - channel for detaching session from topic, buffered.
	// Content is topic name to detach from.
	detach chan string

	// Needed for long polling and grpc.
	lock sync.Mutex

	// Field used only in cluster mode by topic master node.

	// Type of proxy to master request being handled.
	proxyReq ProxyReqType
}

// queueOut attempts to send a ServerComMessage to a session write loop;
// it fails, if the send buffer is full.
func (s *Session) queueOut(msg any) bool {
	if s == nil {
		return true
	}
	if atomic.LoadInt32(&s.terminating) > 0 {
		return true
	}

	select {
	case s.send <- msg:
	default:
		// Never block here since it may also block the topic's run() goroutine.
		logs.Warn("s.queueOut: session's send queue full %s", s.sid)
		return false
	}
	return true
}

// queueOutBytes attempts to send a ServerComMessage already serialized to []byte.
// If the send buffer is full, it fails.
func (s *Session) queueOutBytes(data []byte) bool {
	if s == nil || atomic.LoadInt32(&s.terminating) > 0 {
		return true
	}

	select {
	case s.send <- data:
	default:
		logs.Warn("s.queueOutBytes: session's send queue full %s", s.sid)
		return false
	}
	return true
}

func (s *Session) detachSession(fromTopic string) {
	if atomic.LoadInt32(&s.terminating) == 0 {
		s.detach <- fromTopic
	}
}

func (s *Session) stopSession(data any) {
	s.stop <- data
}

func (s *Session) purgeChannels() {
	for len(s.send) > 0 {
		<-s.send
	}
	for len(s.stop) > 0 {
		<-s.stop
	}
	for len(s.detach) > 0 {
		<-s.detach
	}
}

// cleanUp is called when the session is terminated to perform resource cleanup.
func (s *Session) cleanUp(expired bool) {
	atomic.StoreInt32(&s.terminating, 1)
	s.purgeChannels()
	if !expired {
		s.sessionStoreLock.Lock()
		//globals.sessionStore.Delete(s)
		s.sessionStoreLock.Unlock()
	}

	s.background = false
	s.bkgTimer.Stop()
	// Stop the write loop.
	s.stopSession(nil)
}

// Message received, convert bytes to ClientComMessage and dispatch
func (s *Session) dispatchRaw(raw []byte) {
	var msg types.ServerComMessage

	if atomic.LoadInt32(&s.terminating) > 0 {
		logs.Warn("s.dispatch: message received on a terminating session %s", s.sid)
		return
	}

	if len(raw) == 1 && raw[0] == 0x31 {
		// 0x31 == '1'. This is a network probe message. Respond with a '0':
		s.queueOutBytes([]byte{0x30})
		return
	}

	toLog := raw
	truncated := ""
	if len(raw) > 512 {
		toLog = raw[:512]
		truncated = "<...>"
	}
	logs.Info("in: '%s%s' sid='%s'", toLog, truncated, s.sid)

	if err := json.Unmarshal(raw, &msg); err != nil {
		// Malformed message
		logs.Warn("s.dispatch %s %s", err, s.sid)
		return
	}

	s.dispatch(&msg)
}

func (s *Session) serialize(msg any) (int, any) {
	out, _ := json.Marshal(msg)
	return len(out), out
}

func (s *Session) closeWS() {
	_ = s.ws.Close()
}

func (s *Session) readLoop() {
	defer func() {
		s.closeWS()
		s.cleanUp(false)
	}()

	s.ws.SetReadLimit(defaultMaxMessageSize)
	_ = s.ws.SetReadDeadline(time.Now().Add(pingPeriod))
	s.ws.SetPingHandler(func(string) error {
		_ = s.ws.SetReadDeadline(time.Now().Add(pingPeriod))
		return nil
	})

	for {
		// Read a ClientComMessage
		_, raw, err := s.ws.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure,
				websocket.CloseNormalClosure) {
				logs.Warn("ws: readLoop %s %s", s.sid, err)
			}
			return
		}
		s.dispatchRaw(raw)
	}
}

func (s *Session) sendMessage(msg any) bool {
	if len(s.send) > sendQueueLimit {
		logs.Warn("ws: outbound queue limit exceeded %s", s.sid)
		return false
	}

	data, err := json.Marshal(msg)
	if err != nil {
		logs.Error(err)
		return false
	}

	if err := wsWrite(s.ws, websocket.TextMessage, data); err != nil {
		if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure,
			websocket.CloseNormalClosure) {
			logs.Warn("ws: writeLoop %s %s", s.sid, err)
		}
		return false
	}
	return true
}

func (s *Session) writeLoop() {
	ticker := time.NewTicker(pingPeriod)

	defer func() {
		ticker.Stop()
		// Break readLoop.
		s.closeWS()
	}()

	for {
		select {
		case msg, ok := <-s.send:
			if !ok {
				// Channel closed.
				return
			}
			switch v := msg.(type) {
			default: // serialized message
				if !s.sendMessage(v) {
					return
				}
			}

		case <-s.bkgTimer.C:
			if s.background {
				s.background = false
			}

		case msg := <-s.stop:
			// Shutdown requested, don't care if the message is delivered
			if msg != nil {
				_ = wsWrite(s.ws, websocket.TextMessage, msg)
			}
			return

		case <-ticker.C:
			if err := wsWrite(s.ws, websocket.PongMessage, nil); err != nil {
				if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure,
					websocket.CloseNormalClosure) {
					logs.Warn("ws: writeLoop pong %s %s", s.sid, err)
				}
				return
			}
		}
	}
}

// Writes a message with the given message type (mt) and payload.
func wsWrite(ws *websocket.Conn, mt int, msg any) error {
	var bits []byte
	if msg != nil {
		bits = msg.([]byte)
	} else {
		bits = []byte{}
	}
	_ = ws.SetWriteDeadline(time.Now().Add(writeWait))
	return ws.WriteMessage(mt, bits)
}
