package postgres

import (
	"errors"

	"github.com/lfourky/go-rest-service-template/pkg/model/domain"
	"gorm.io/gorm"
)

type User struct {
	db *gorm.DB
}

func (p *User) FindByID(id domain.UUID) (*domain.User, error) {
	var user domain.User

	err := p.db.Where("id = ?", id).First(&user).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}

	return &user, err
}

func (p *User) FindByEmail(email string) (*domain.User, error) {
	var user domain.User

	err := p.db.Where("email = ?", email).First(&user).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}

	return &user, err
}

func (p *User) Create(user *domain.User) error {
	return p.db.Create(user).Error
}

func (p *User) FindAll() ([]*domain.User, error) {
	var users []*domain.User
	return users, p.db.Preload("Items").Find(&users).Error
}
