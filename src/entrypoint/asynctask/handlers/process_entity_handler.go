package handlers

import (
	"context"
	"example.com/PROJECT_NAME/src/core/domain/ports/async"
	"example.com/PROJECT_NAME/src/core/utils"
)

type ProcessEntityHandler struct {
	// TODO: use case dependency
}

// @inject
func NewProcessEntityHandler() *ProcessEntityHandler {
	return &ProcessEntityHandler{}
}

func (this *ProcessEntityHandler) Handle(ctx context.Context, task *async.Task) error {
	entityID, err := utils.JsonUnmarshal[uint](task.Payload)
	if err != nil {
		return err
	}
	_ = entityID
	// TODO: this.useCase.Invoke(ctx, entityID)
	return nil
}
