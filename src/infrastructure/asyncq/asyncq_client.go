package async

import "github.com/hibiken/asynq"

// @inject
func NewAsynqClient(redis asynq.RedisClientOpt) *asynq.Client {
	return asynq.NewClient(redis)
}
