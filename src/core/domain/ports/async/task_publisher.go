package async

import "time"

type Option struct {
	ProcessIn     *time.Duration
	ProcessAt     *time.Time
	MaxRetryCount int
	TimeOut       *time.Duration
}

type TaskPublisher interface {
	Publish(task *Task, option Option) error
}
