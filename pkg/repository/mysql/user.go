package mysql

import (
	"errors"

	"github.com/lfourky/go-transaction-management/pkg/model"
)

type User struct {
	*Repository
}

// This method demonstrates how gorm.DB is used in this architecture.
func (u *User) FindUserByID(id string) (*model.User, error) {
	var user *model.User
	err := u.db.Find(&user).Where("id = ?", id).Error
	return user, err
}

func (u *User) FindUserByName(name string) (*model.User, error) {
	return nil, errors.New("not implemented")
}
