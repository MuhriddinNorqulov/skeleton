package echohttp

import (
	"example.com/PROJECT_NAME/src/core/domain/ports/httpport"
	"example.com/PROJECT_NAME/src/infrastructure/echohttp/defaults"
	"example.com/PROJECT_NAME/src/infrastructure/echohttp/mapper"
	"example.com/PROJECT_NAME/src/infrastructure/env"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

// @inject
func NewEcho(validator *defaults.RequestValidator) *echo.Echo {
	e := echo.New()
	e.Validator = validator
	return e
}

type EchoServerImpl struct {
	echo                    *echo.Echo
	env                     *env.Env
	recoveryMiddleware      *defaults.RecoveryMiddleware
	httpLoggerMiddleware    *defaults.HttpLoggerMiddleware
	consoleLoggerMiddleware *defaults.ConsoleLoggerMiddleware
}

// @inject
func NewEchoServerImpl(
	echo *echo.Echo,
	cfg *env.Env,
	recoveryMiddleware *defaults.RecoveryMiddleware,
	httpLoggerMiddleware *defaults.HttpLoggerMiddleware,
	consoleLoggerMiddleware *defaults.ConsoleLoggerMiddleware,
) httpport.HTTPServer {
	return &EchoServerImpl{
		echo:                    echo,
		env:                     cfg,
		recoveryMiddleware:      recoveryMiddleware,
		httpLoggerMiddleware:    httpLoggerMiddleware,
		consoleLoggerMiddleware: consoleLoggerMiddleware,
	}
}

func (this *EchoServerImpl) Use(middlewares ...httpport.Middleware) {
	_mws := mapper.MiddlewareListMapper(middlewares)
	this.echo.Use(_mws...)
}

func (this *EchoServerImpl) Run() {
	this.echo.Logger.Fatal(this.echo.Start(this.env.HttpAddress))
}

func (this *EchoServerImpl) Group(prefix string, middlewares ...httpport.Middleware) httpport.Group {
	_mws := mapper.MiddlewareListMapper(middlewares)
	group := this.echo.Group(prefix, _mws...)
	return NewEchoGroupImpl(group)
}

func (this *EchoServerImpl) Init() {
	this.echo.Use(this.consoleLoggerMiddleware.Wrap)

	this.echo.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins:     this.env.CorsAllowOrigin,
		AllowHeaders:     []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept, echo.HeaderAuthorization, "Idempotency-Key"},
		AllowCredentials: this.env.CorsAllowCredentials,
	}))

	this.echo.Use(this.httpLoggerMiddleware.Wrap)
	this.echo.Use(this.recoveryMiddleware.Wrap)
	this.echo.Use(defaults.ContextMiddleware)
}
