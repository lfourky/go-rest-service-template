package dto

type User struct {
	ID    string  `json:"id"`
	Name  string  `json:"name"`
	Email string  `json:"email"`
	Items []*Item `json:"items"`
}

type RegisterUserRequest struct {
	Email string `json:"email" validate:"required"`
	Name  string `json:"name" validate:"required"`
}

type RegisterUserResponse struct {
	User
}
