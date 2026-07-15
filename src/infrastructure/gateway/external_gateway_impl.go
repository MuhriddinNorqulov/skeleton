package gateway

import (
	"context"
	"fmt"

	"example.com/PROJECT_NAME/src/core/application/response"
	gatewayport "example.com/PROJECT_NAME/src/core/domain/ports/gateway"
	"example.com/PROJECT_NAME/src/core/utils"
	"example.com/PROJECT_NAME/src/infrastructure/env"
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
