package gateway

import (
	"context"
	"fmt"

	"github.com/muhriddinnorqulov/skeleton/src/core/application/response"
	gatewayport "github.com/muhriddinnorqulov/skeleton/src/core/domain/ports/gateway"
	"github.com/muhriddinnorqulov/skeleton/src/core/utils"
	"github.com/muhriddinnorqulov/skeleton/src/infrastructure/env"
)

type ExternalGatewayImpl struct {
	env *env.Env
}

// @inject
func NewExternalGatewayImpl(env *env.Env) gatewayport.ExternalServiceGateway {
	return &ExternalGatewayImpl{env: env}
}

func (this *ExternalGatewayImpl) DoSomething(ctx context.Context, input string) (string, error) {
	// TODO: tashqi API chaqiruvi
	return "", response.NewSafeError(response.CodeGatewayError,
		fmt.Errorf("not implemented"), utils.CallerPath(1))
}
