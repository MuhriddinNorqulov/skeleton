package tasks

import (
	"github.com/muhriddinnorqulov/skeleton/src/core/domain/entity/enum"
	"github.com/muhriddinnorqulov/skeleton/src/core/domain/ports/async"
	"github.com/muhriddinnorqulov/skeleton/src/core/utils"
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
