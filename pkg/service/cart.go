package service

import (
	"github.com/lfourky/db-access/pkg/model"
	"github.com/lfourky/db-access/pkg/repository"
)

type Cart struct {
	store repository.Store
}

func NewCart(s repository.Store) *Cart {
	return &Cart{
		store: s,
	}
}

func (c *Cart) DoSomethingCartRelated() (*model.Cart, error) {
	cart, err := c.store.Carts().FindCartByID("some_cart_id")
	if err != nil {
		return nil, err
	}

	return cart, nil
}

func (c *Cart) DemonstrateTransaction() (*model.User, error) {
	tx, err := c.store.BeginTransaction()
	if err != nil {
		return nil, err
	}

	user, err := tx.Users().FindUserByID("some_user_id")
	if err != nil {
		_ = tx.Rollback()
		return nil, err
	}

	cart, err := tx.Carts().FindCartByID("some_cart_id")
	if err != nil {
		_ = tx.Rollback
		return nil, err
	}

	user.ShoppingCart = cart

	item, err := tx.Items().FindItemByID("some_item_id")
	if err != nil {
		_ = tx.Rollback()
		return nil, err
	}

	cart.Items = append(cart.Items, item)

	if err = tx.Commit(); err != nil {
		return nil, err
	}

	return user, nil
}
