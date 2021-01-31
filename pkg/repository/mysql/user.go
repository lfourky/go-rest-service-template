package mysql

import (
	"github.com/lfourky/go-rest-service-template/pkg/model"
)

type User struct {
	*Repository
}

func (u *User) FindByID(id string) (*model.User, error) {
	var user *model.User
	err := u.db.Find(&user).Where("id = ?", id).Error
	return user, err
}

func (u *User) FindByEmail(email string) (*model.User, error) {
	var user *model.User
	err := u.db.Find(&user).Where("email = ?", email).Error
	return user, err
}

func (u *User) Create(user *model.User) error {
	return u.db.Create(user).Error
}
