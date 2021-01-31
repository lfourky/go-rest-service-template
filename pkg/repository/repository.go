package repository

import (
	"github.com/lfourky/go-rest-service-template/pkg/model"
)

type Store interface {
	Users() User
	Items() Item
	UnitOfWork
}

type UnitOfWork interface {
	BeginTransaction() (Store, error)
	Commit() error
	Rollback() error
}

type User interface {
	Create(user *model.User) error
	FindByEmail(email string) (*model.User, error)
	FindByID(id string) (*model.User, error)
}

type Item interface {
	Create(item *model.Item) error
	Delete(item *model.Item) error
}
