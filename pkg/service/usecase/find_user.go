package usecase

import (
	"github.com/lfourky/go-transaction-management/pkg/model"
	"github.com/lfourky/go-transaction-management/pkg/repository"
)

type FindUser struct {
	store repository.Store
}

func (u *FindUser) FindUserByID(id string) (*model.User, error) {
	return u.store.Users().FindByID(id)
}

func (u *FindUser) FindUserByEmail(email string) (*model.User, error) {
	return u.store.Users().FindByEmail(email)
}
