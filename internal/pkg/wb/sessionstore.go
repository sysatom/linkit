package wb

import (
	"container/list"
	"github.com/google/uuid"
	"github.com/gorilla/websocket"
	"github.com/sysatom/linkit/internal/pkg/logs"
	"sync"
	"time"
)

// WaitGroup with a semaphore functionality
// (limiting number of threads/goroutines accessing the guarded resource simultaneously).
type boundedWaitGroup struct {
	wg  sync.WaitGroup
	sem chan struct{}
}

func newBoundedWaitGroup(capacity int) *boundedWaitGroup {
	return &boundedWaitGroup{sem: make(chan struct{}, capacity)}
}

func (w *boundedWaitGroup) Add(delta int) {
	if delta <= 0 {
		return
	}
	for i := 0; i < delta; i++ {
		w.sem <- struct{}{}
	}
	w.wg.Add(delta)
}

func (w *boundedWaitGroup) Done() {
	select {
	case _, ok := <-w.sem:
		if !ok {
			logs.Warn("boundedWaitGroup.sem closed.")
		}
	default:
		logs.Warn("boundedWaitGroup.Done() called before Add().")
	}
	w.wg.Done()
}

func (w *boundedWaitGroup) Wait() {
	w.wg.Wait()
}

// SessionStore holds live sessions. Long polling sessions are stored in a linked list with
// most recent sessions on top. In addition, all sessions are stored in a map indexed by session ID.
type SessionStore struct {
	lock sync.Mutex

	// Support for long polling sessions: a list of sessions sorted by last access time.
	// Needed for cleaning abandoned sessions.
	lru      *list.List
	lifeTime time.Duration

	// All sessions indexed by session ID
	sessCache map[string]*Session
}

// NewSession creates a new session and saves it to the session store.
func (ss *SessionStore) NewSession(conn any, sid string) (*Session, int) {
	var s Session

	if sid == "" {
		s.sid = uuid.New().String()
	} else {
		s.sid = sid
	}

	ss.lock.Lock()
	if _, found := ss.sessCache[s.sid]; found {
		logs.Warn("ERROR! duplicate session ID %s", s.sid)
	}
	ss.lock.Unlock()

	switch c := conn.(type) {
	case *websocket.Conn:
		s.ws = c
	default:
		logs.Warn("session: unknown connection type %s", conn)
	}

	s.send = make(chan any, sendQueueLimit+32) // buffered
	s.stop = make(chan any, 1)                 // Buffered by 1 just to make it non-blocking
	s.detach = make(chan string, 64)           // buffered

	s.bkgTimer = time.NewTimer(time.Hour)
	s.bkgTimer.Stop()

	s.lastTouched = time.Now()

	ss.lock.Lock()

	ss.sessCache[s.sid] = &s

	// Expire stale long polling sessions: ss.lru contains only long polling sessions.
	// If ss.lru is empty this is a noop.
	var expired []*Session
	expire := s.lastTouched.Add(-ss.lifeTime)
	for elem := ss.lru.Back(); elem != nil; elem = ss.lru.Back() {
		sess := elem.Value.(*Session)
		if sess.lastTouched.Before(expire) {
			ss.lru.Remove(elem)
			delete(ss.sessCache, sess.sid)
			expired = append(expired, sess)
		} else {
			break // don't need to traverse further
		}
	}

	numSessions := len(ss.sessCache)

	ss.lock.Unlock()

	// Deleting long polling sessions.
	for _, sess := range expired {
		// This locks the session. Thus cleaning up outside the
		// sessionStore lock. Otherwise, deadlock.
		sess.cleanUp(true)
	}

	return &s, numSessions
}

// Get fetches a session from store by session ID.
func (ss *SessionStore) Get(sid string) *Session {
	ss.lock.Lock()
	defer ss.lock.Unlock()

	if sess := ss.sessCache[sid]; sess != nil {
		return sess
	}

	return nil
}

// Delete removes session from store.
func (ss *SessionStore) Delete(s *Session) {
	ss.lock.Lock()
	defer ss.lock.Unlock()

	delete(ss.sessCache, s.sid)
}

// Range calls given function for all sessions. It stops if the function returns false.
func (ss *SessionStore) Range(f func(sid string, s *Session) bool) {
	ss.lock.Lock()
	for sid, s := range ss.sessCache {
		if !f(sid, s) {
			break
		}
	}
	ss.lock.Unlock()
}

// Shutdown terminates sessionStore. No need to clean up.
// Don't send to clustered sessions, their servers are not being shut down.
func (ss *SessionStore) Shutdown() {
	ss.lock.Lock()
	defer ss.lock.Unlock()

	shutdown := time.Now().String()
	for _, s := range ss.sessCache {
		_, data := s.serialize(shutdown)
		s.stopSession(data)
	}

	logs.Info("SessionStore shut down, sessions terminated: %s", len(ss.sessCache))
}

// NewSessionStore initializes a session store.
func NewSessionStore(lifetime time.Duration) *SessionStore {
	ss := &SessionStore{
		lru:      list.New(),
		lifeTime: lifetime,

		sessCache: make(map[string]*Session),
	}

	return ss
}
