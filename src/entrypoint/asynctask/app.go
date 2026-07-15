package asynctask

import (
	"example.com/PROJECT_NAME/src/core/domain/entity/enum"
	"example.com/PROJECT_NAME/src/core/domain/ports/async"
	"example.com/PROJECT_NAME/src/entrypoint/asynctask/handlers"
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
