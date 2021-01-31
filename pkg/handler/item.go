package handler

import "github.com/lfourky/go-rest-service-template/pkg/service"

type Item struct {
	itemService service.Item
}

func NewItemHandler(itemService service.Item) *Item {
	return &Item{
		itemService: itemService,
	}
}
