package handler

import (
	"github.com/labstack/echo/v4"
)

func bindAndValidate(request interface{}, ctx echo.Context) error {
	if err := ctx.Bind(request); err != nil {
		return err
	}

	if err := ctx.Validate(request); err != nil {
		return err
	}

	return nil
}
