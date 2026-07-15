package async

import (
	"context"
	"example.com/PROJECT_NAME/src/core/domain/entity/enum"
	"example.com/PROJECT_NAME/src/core/domain/ports/async"
	"example.com/PROJECT_NAME/src/infrastructure/asyncq/mapper"

	"github.com/hibiken/asynq"
)

type AsynqServerImpl struct {
	s   *asynq.Server
	mux *asynq.ServeMux
}

// @inject
func NewAsynqServerImpl(s *asynq.Server, mux *asynq.ServeMux) async.AsyncServer {
	return &AsynqServerImpl{s: s, mux: mux}
}

func (this *AsynqServerImpl) Init() {}

func (this *AsynqServerImpl) Run() error {
	return this.s.Run(this.mux)
}

func (this *AsynqServerImpl) Use(middlewares ...async.AsyncTaskMiddleware) {
	for _, middleware := range middlewares {
		this.mux.Use(mapper.AsynqTaskMiddlewareMapper(middleware))
	}
}

func (this *AsynqServerImpl) HandlerFunc(taskName enum.TaskType, handler async.AsyncTaskHandler) {
	this.mux.HandleFunc(string(taskName), mapper.ToAsynQTask(handler).ProcessTask)
}

func (this *AsynqServerImpl) ProcessTask(ctx context.Context, task *async.Task) error {
	return this.mux.ProcessTask(ctx, asynq.NewTask(string(task.TaskType), task.Payload))
}
