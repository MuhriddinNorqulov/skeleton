package gorillaws

import (
	"net/http"

	"github.com/gorilla/websocket"
)

type GorillaUpgrader struct {
	u websocket.Upgrader
}

// @inject
func NewGorillaUpgrader() *GorillaUpgrader {
	return &GorillaUpgrader{
		u: websocket.Upgrader{
			CheckOrigin: func(r *http.Request) bool { return true },
		},
	}
}

func (this *GorillaUpgrader) Upgrade(w http.ResponseWriter, r *http.Request) (*websocket.Conn, error) {
	return this.u.Upgrade(w, r, nil)
}
