package dto

type Item struct {
	ID     string `json:"id"`
	Name   string `json:"name"`
	UserID string `json:"user_id,omitempty"`
}

type CreateItemRequest struct {
	UserID string `json:"user_id" validate:"required"`
	Name   string `json:"name" validate:"required"`
}

type CreateItemResponse struct {
	Item
}
