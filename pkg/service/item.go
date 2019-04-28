package service

import (
	"github.com/lfourky/db-access/pkg/model"
	"github.com/lfourky/db-access/pkg/repository"
)

type Item struct {
	store repository.Store
}

func NewItem(s repository.Store) *Item {
	return &Item{
		store: s,
	}
}

func (i *Item) DoSomethingItemRelated() (*model.Item, error) {
	item, err := i.store.Items().FindItemByID("some_item_id")
	if err != nil {
		return nil, err
	}

	return item, err
}

func (i *Item) DemonstrateTransaction() (*model.Item, error) {
	tx, err := i.store.BeginTransaction()
	if err != nil {
		return nil, err
	}

	cart, err := tx.Carts().FindCartByID("some_cart_id")
	if err != nil {
		_ = tx.Rollback()
		return nil, err
	}

	item, err := tx.Items().FindItemByID("some_item_id")
	if err != nil {
		_ = tx.Rollback()
		return nil, err
	}

	cart.Items = append(cart.Items, item)

	if err = tx.Commit(); err != nil {
		return nil, err
	}

	return item, nil
}
