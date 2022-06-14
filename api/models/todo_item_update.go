package models

type TodoIteUpdateRequest struct {
	Done bool `json:"done" validate:"required"`
}
