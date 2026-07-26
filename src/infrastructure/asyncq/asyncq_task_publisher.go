package async

import (
	"github.com/muhriddinnorqulov/skeleton/src/core/domain/ports/async"
	"strconv"
	"time"

	"github.com/hibiken/asynq"
)

type AsynqTaskPublisher struct {
	client *asynq.Client
}

// @inject
func NewAsynqTaskPublisher(client *asynq.Client) async.TaskPublisher {
	return &AsynqTaskPublisher{client: client}
}

func (this *AsynqTaskPublisher) Publish(task *async.Task, opt async.Option) error {
	now := time.Now().UnixNano()
	taskID := strconv.FormatInt(now, 10)

	asynqTask := asynq.NewTask(string(task.TaskType), task.Payload)

	var options = []asynq.Option{
		asynq.MaxRetry(opt.MaxRetryCount),
		asynq.TaskID(taskID),
	}

	if t := opt.ProcessIn; t != nil {
		options = append(options, asynq.ProcessIn(*t))
	}
	if t := opt.ProcessAt; t != nil {
		options = append(options, asynq.ProcessAt(*t))
	}
	if t := opt.TimeOut; t != nil {
		options = append(options, asynq.Timeout(*t))
	}

	_, err := this.client.Enqueue(asynqTask, options...)
	return err
}
