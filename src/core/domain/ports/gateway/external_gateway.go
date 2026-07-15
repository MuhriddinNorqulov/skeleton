package gateway

import "context"

type ExternalServiceGateway interface {
	DoSomething(ctx context.Context, input string) (string, error)
}
