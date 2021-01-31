package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/lfourky/go-rest-service-template/pkg/model/dto"
	"github.com/lfourky/go-rest-service-template/pkg/server"
	"github.com/lfourky/go-rest-service-template/pkg/usecase"
)

type User struct {
	users usecase.User
}

func NewUser(users usecase.User) *User {
	return &User{
		users: users,
	}
}

func (h *User) RegisterUser(ctx echo.Context) error {
	req := &dto.RegisterUserRequest{}
	if err := bindAndValidate(req, ctx); err != nil {
		return server.NewHTTPError(err)
	}

	resp, err := h.users.RegisterUser(req)
	if err != nil {
		return server.NewHTTPError(err)
	}

	return ctx.JSON(http.StatusCreated, resp)
}

func (h *User) FindAll(ctx echo.Context) error {
	resp, err := h.users.FindAllUsers()
	if err != nil {
		return server.NewHTTPError(err)
	}

	return ctx.JSON(http.StatusOK, resp)
}
