package handler

import "github.com/lfourky/go-rest-service-template/pkg/service"

type User struct {
	userService service.User
}

func NewUserHandler(userService service.User) *User {
	return &User{
		userService: userService,
	}
}
