package repository

import "github.com/lfourky/db-access/pkg/model"

type Store interface {
	Users() User
	Items() Item
	Carts() Cart
	UnitOfWork
}

type UnitOfWork interface {
	BeginTransaction() (Store, error)
	Commit() error
	Rollback() error
}

type User interface {
	FindUserByID(id string) (*model.User, error)
	FindUserByName(name string) (*model.User, error)
}

type Item interface {
	FindItemByID(id string) (*model.Item, error)
	FindItemByName(name string) (*model.Item, error)
}

type Cart interface {
	FindCartByID(id string) (*model.Cart, error)
	FindCartByItem(i model.Item) (*model.Cart, error)
}
