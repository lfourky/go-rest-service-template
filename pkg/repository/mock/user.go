package mock

import (
	"github.com/lfourky/go-rest-service-template/pkg/model/domain"
	"github.com/stretchr/testify/mock"
)

type User struct {
	mock.Mock
}

func (m *User) Create(user *domain.User) error {
	args := m.Called(user)

	return args.Error(0)
}

func (m *User) FindAll() ([]*domain.User, error) {
	args := m.Called()

	if args.Get(0) != nil {
		return args.Get(0).([]*domain.User), args.Error(1)
	}

	return nil, args.Error(1)
}

func (m *User) FindByID(id domain.UUID) (*domain.User, error) {
	args := m.Called(id)

	if args.Get(0) != nil {
		return args.Get(0).(*domain.User), args.Error(1)
	}

	return nil, args.Error(1)
}

func (m *User) FindByEmail(email string) (*domain.User, error) {
	args := m.Called(email)

	if args.Get(0) != nil {
		return args.Get(0).(*domain.User), args.Error(1)
	}

	return nil, args.Error(1)
}
