package wb

import (
	"github.com/gorilla/websocket"
	"github.com/sysatom/linkit/internal/pkg/logs"
	"github.com/sysatom/linkit/internal/pkg/setting"
	"net/http"
	"net/url"
	"time"
)

var sessionStore = NewSessionStore(idleSessionTimeout + 15*time.Second)

func Init() {
	u := url.URL{
		Scheme: "ws",
		Host:   setting.Get().ServerHost,
		Path:   "/extra/session",
	}
	logs.Info("connecting to %s", u.String())

	header := http.Header{}
	header.Set("X-AccessToken", setting.Get().AccessToken)

	conn, _, err := websocket.DefaultDialer.Dial(u.String(), header)
	if err != nil {
		logs.Error(err)
		return
	}

	sess, count := sessionStore.NewSession(conn, "")
	logs.Info("ws: session started %s %d", sess.sid, count)

	//go func() {
	//	for {
	//		_, message, err := conn.ReadMessage()
	//		if err != nil {
	//			logs.Info("read: %s", err)
	//			return
	//		}
	//		logs.Info("recv: %s", message)
	//	}
	//}()

	// Do work in goroutines to return from serveWebSocket() to release file pointers.
	// Otherwise, "too many open files" will happen.
	go sess.writeLoop()
	go sess.readLoop()
}
