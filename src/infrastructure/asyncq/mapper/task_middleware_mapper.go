package mapper

import (
	"context"
	"example.com/PROJECT_NAME/src/core/domain/ports/async"

	"github.com/hibiken/asynq"
)

func AsynqTaskMiddlewareMapper(middleware async.AsyncTaskMiddleware) asynq.MiddlewareFunc {
	return func(next asynq.Handler) asynq.Handler {
		wrapped := middleware(async.AsyncTaskHandlerFunc(func(ctx context.Context, task *async.Task) error {
			return next.ProcessTask(ctx, asynq.NewTask(string(task.TaskType), task.Payload))
		}))
		return asynq.HandlerFunc(func(ctx context.Context, t *asynq.Task) error {
			return wrapped.ProcessTask(ctx, &async.Task{Payload: t.Payload()})
		})
	}
}
