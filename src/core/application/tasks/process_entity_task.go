package tasks

import (
	"example.com/PROJECT_NAME/src/core/domain/entity/enum"
	"example.com/PROJECT_NAME/src/core/domain/ports/async"
	"example.com/PROJECT_NAME/src/core/utils"
	"time"
)

type ProcessEntityTask struct {
	pub async.TaskPublisher
}

// @inject
func NewProcessEntityTask(pub async.TaskPublisher) *ProcessEntityTask {
	return &ProcessEntityTask{pub: pub}
}

func (this *ProcessEntityTask) Publish(entityID uint) error {
	payload, _ := utils.JsonMarshal(entityID)
	task := async.NewTask(enum.TaskProcessEntity, payload)
	delay := 1 * time.Minute
	return this.pub.Publish(task, async.Option{ProcessIn: &delay, MaxRetryCount: 5})
}
