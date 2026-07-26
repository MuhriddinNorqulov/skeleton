package mapper

import (
	"context"
	"github.com/muhriddinnorqulov/skeleton/src/core/domain/ports/async"

	"github.com/hibiken/asynq"
)

type asynqTaskAdapter struct {
	handler async.AsyncTaskHandler
}

func ToAsynQTask(handler async.AsyncTaskHandler) *asynqTaskAdapter {
	return &asynqTaskAdapter{handler: handler}
}

func (this *asynqTaskAdapter) ProcessTask(ctx context.Context, t *asynq.Task) error {
	task := &async.Task{
		TaskType: "",
		Payload:  t.Payload(),
	}
	return this.handler.ProcessTask(ctx, task)
}
