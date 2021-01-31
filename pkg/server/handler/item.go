package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/lfourky/go-rest-service-template/pkg/model/dto"
	"github.com/lfourky/go-rest-service-template/pkg/server"
	"github.com/lfourky/go-rest-service-template/pkg/usecase"
)

type Item struct {
	items usecase.Item
}

func NewItem(items usecase.Item) *Item {
	return &Item{
		items: items,
	}
}

func (h *Item) CreateItem(ctx echo.Context) error {
	req := &dto.CreateItemRequest{}
	if err := bindAndValidate(req, ctx); err != nil {
		return server.NewHTTPError(err)
	}

	resp, err := h.items.CreateItem(req)
	if err != nil {
		return server.NewHTTPError(err)
	}

	return ctx.JSON(http.StatusCreated, resp)
}

func (h *Item) FindAll(ctx echo.Context) error {
	resp, err := h.items.FindAllItems()
	if err != nil {
		return server.NewHTTPError(err)
	}

	return ctx.JSON(http.StatusOK, resp)
}
