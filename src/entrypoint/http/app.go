package http

import (
	"github.com/muhriddinnorqulov/skeleton/src/core/domain/ports/httpport"
	"github.com/muhriddinnorqulov/skeleton/src/core/domain/ports/wsport"
	"github.com/muhriddinnorqulov/skeleton/src/entrypoint/http/groups"
	"github.com/muhriddinnorqulov/skeleton/src/entrypoint/http/interceptor/middlewares"
	"github.com/muhriddinnorqulov/skeleton/src/infrastructure/env"
)

type App struct {
	server httpport.HTTPServer
	env    *env.Env

	jwtAuthMiddleware  *middlewares.JwtAuthMiddleware
	responseMiddleware *middlewares.ResponseMiddleware

	entityGroup *groups.EntityGroup
	wsGroup     *groups.WsGroup

	ws wsport.Ws
}

// @inject
func NewApp(
	server httpport.HTTPServer,
	env *env.Env,
	jwtAuthMiddleware *middlewares.JwtAuthMiddleware,
	responseMiddleware *middlewares.ResponseMiddleware,
	entityGroup *groups.EntityGroup,
	wsGroup *groups.WsGroup,
	ws wsport.Ws,
) *App {
	return &App{
		server:             server,
		env:                env,
		jwtAuthMiddleware:  jwtAuthMiddleware,
		responseMiddleware: responseMiddleware,
		entityGroup:        entityGroup,
		wsGroup:            wsGroup,
		ws:                 ws,
	}
}

func (this *App) Init() {
	this.server.Init()
	this.initMiddlewares()
	this.initGroups()
}

func (this *App) initMiddlewares() {
	this.server.Use(this.jwtAuthMiddleware.Wrap)
	this.server.Use(this.responseMiddleware.Wrap)
}

func (this *App) initGroups() {
	v1 := this.group("/v1")

	this.entityGroup.RegisterRoutes(v1.Group("/entity"))
	this.wsGroup.RegisterRoutes(v1.Group("/ws"))
}

func (this *App) Start() {
	this.server.Run()
}

func (this *App) group(prefix string, mws ...httpport.Middleware) httpport.Group {
	basePath := "/api" + prefix
	return this.server.Group(basePath, mws...)
}
