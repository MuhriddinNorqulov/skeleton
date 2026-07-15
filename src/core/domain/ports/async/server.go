package async

import (
	"context"
	"example.com/PROJECT_NAME/src/core/domain/entity/enum"
)

type AsyncServer interface {
	Init()
	Run() error
	Use(middlewares ...AsyncTaskMiddleware)
	HandlerFunc(taskName enum.TaskType, handler AsyncTaskHandler)
	ProcessTask(ctx context.Context, task *Task) error
}
