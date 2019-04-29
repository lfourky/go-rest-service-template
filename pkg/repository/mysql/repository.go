package mysql

import (
	"github.com/jinzhu/gorm"
	"github.com/lfourky/go-transaction-management/pkg/repository"
)

type Repository struct {
	db    *gorm.DB
	users repository.User
	items repository.Item
	carts repository.Cart
}

func NewRepository(db *gorm.DB) *Repository {
	r := &Repository{
		db: db,
	}

	r.users = &User{Repository: r}
	r.items = &Item{Repository: r}
	r.carts = &Cart{Repository: r}

	return r
}

func (r *Repository) BeginTransaction() (repository.Store, error) {
	tx := r.db.Begin()
	if tx.Error != nil {
		return nil, tx.Error
	}
	return NewRepository(tx), nil
}

func (r *Repository) Commit() error {
	return r.db.Commit().Error
}

func (r *Repository) Rollback() error {
	return r.db.Rollback().Error
}

func (r *Repository) Users() repository.User {
	return r.users
}

func (r *Repository) Items() repository.Item {
	return r.items
}

func (r *Repository) Carts() repository.Cart {
	return r.carts
}
