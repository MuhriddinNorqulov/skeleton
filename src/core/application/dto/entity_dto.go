package dto

type CreateEntityRequest struct {
	Name string `json:"name" validate:"required,min=3"`
}

type CreateEntityResult struct {
	ID uint `json:"id"`
}
