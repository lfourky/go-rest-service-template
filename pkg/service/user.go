package service

import (
	"github.com/lfourky/db-access/pkg/model"
	"github.com/lfourky/db-access/pkg/repository"
)

type User struct {
	store repository.Store
}

func NewUser(s repository.Store) *User {
	return &User{
		store: s,
	}
}

func (u *User) DoSomethingUserRelated() (*model.User, error) {
	user, err := u.store.Users().FindUserByID("some_user_id")
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (u *User) DemonstrateTransaction() (*model.Item, error) {
	tx, err := u.store.BeginTransaction()
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
