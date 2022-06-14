package models

type TodoItemCreateRequest struct {
	Description string `json:"description" validate:"required"`
}
