package handler

import "github.com/lfourky/go-transaction-management/pkg/service"

type Item struct {
	itemService service.Item
}

func NewItemHandler(itemService service.Item) *Item {
	return &Item{
		itemService: itemService,
	}
}
