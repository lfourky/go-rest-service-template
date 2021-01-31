package usecase

import (
	"github.com/lfourky/go-rest-service-template/pkg/model"
	"github.com/lfourky/go-rest-service-template/pkg/repository"
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
