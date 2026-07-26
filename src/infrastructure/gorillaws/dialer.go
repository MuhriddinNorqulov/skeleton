package gorillaws

import (
	"time"

	"github.com/muhriddinnorqulov/skeleton/src/core/domain/ports/config"

	"github.com/gorilla/websocket"
)

type Dialer struct {
	cfg config.ConfigProvider
}

// @inject
func NewDialer(cfg config.ConfigProvider) *Dialer {
	return &Dialer{cfg: cfg}
}

func (this *Dialer) Dial(url string) (*websocket.Conn, error) {
	ws, resp, err := websocket.DefaultDialer.Dial(url, nil)
	if resp != nil && resp.Body != nil {
		_ = resp.Body.Close()
	}
	return ws, err
}

func (this *Dialer) BackoffDelay(attempt int) time.Duration {
	base := this.cfg.GetWsReconnectBaseDelay()
	max := this.cfg.GetWsReconnectMaxDelay()
	delay := base * time.Duration(attempt+1)
	if delay > max {
		return max
	}
	return delay
}
