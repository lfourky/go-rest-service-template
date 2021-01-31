package repository

import (
	"github.com/lfourky/go-rest-service-template/pkg/model/domain"
)

type Store interface {
	User() User
	Item() Item
	UnitOfWork
}

type UnitOfWork interface {
	BeginTransaction() (Store, error)
	Commit() error
	Rollback() error
}

type User interface {
	Create(user *domain.User) error
	FindAll() ([]*domain.User, error)
	FindByID(id domain.UUID) (*domain.User, error)
	FindByEmail(email string) (*domain.User, error)
}

type Item interface {
	Create(item *domain.Item) error
	FindAll() ([]*domain.Item, error)
}
