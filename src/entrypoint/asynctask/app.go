package asynctask

import (
	"github.com/muhriddinnorqulov/skeleton/src/core/domain/entity/enum"
	"github.com/muhriddinnorqulov/skeleton/src/core/domain/ports/async"
	"github.com/muhriddinnorqulov/skeleton/src/entrypoint/asynctask/handlers"
)

type App struct {
	server async.AsyncServer

	processEntityHandler *handlers.ProcessEntityHandler
}

// @inject
func NewAsyncApp(
	server async.AsyncServer,
	processEntityHandler *handlers.ProcessEntityHandler,
) *App {
	return &App{
		server:               server,
		processEntityHandler: processEntityHandler,
	}
}

func (this *App) Init() {
	this.server.Init()
	this.initHandlers()
}

func (this *App) initHandlers() {
	this.server.HandlerFunc(enum.TaskProcessEntity, async.AsyncTaskHandlerFunc(this.processEntityHandler.Handle))
}

func (this *App) Start() {
	if err := this.server.Run(); err != nil {
		panic(err)
	}
}
