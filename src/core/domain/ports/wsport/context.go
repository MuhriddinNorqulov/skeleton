package wsport

import (
	"context"
	"time"
)

type Context interface {
	context.Context
	SetReadDeadline(d time.Duration)
}
