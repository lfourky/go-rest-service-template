package mysql

import (
	"errors"

	"github.com/lfourky/go-rest-service-template/pkg/model/domain"
	"gorm.io/gorm"
)

type User struct {
	db *gorm.DB
}

func (m *User) FindByID(id domain.UUID) (*domain.User, error) {
	var user domain.User

	err := m.db.Find(&user).Where("id = ?", id).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}

	return &user, err
}

func (m *User) FindByEmail(email string) (*domain.User, error) {
	var user domain.User

	err := m.db.Where("email = ?", email).First(&user).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}

	return &user, err
}

func (m *User) Create(user *domain.User) error {
	return m.db.Create(user).Error
}

func (m *User) FindAll() ([]*domain.User, error) {
	var users []*domain.User
	return users, m.db.Find(&users).Error
}
