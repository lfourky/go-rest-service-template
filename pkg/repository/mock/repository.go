package mock

import (
	"github.com/lfourky/go-rest-service-template/pkg/repository"
	"github.com/stretchr/testify/mock"
)

type Store struct {
	mock.Mock
	ItemMock *Item
	UserMock *User
}

func New() *Store {
	return &Store{
		UserMock: &User{},
		ItemMock: &Item{},
	}
}

func (s *Store) BeginTransaction() (repository.Store, error) {
	args := s.Called()

	if args.Get(0) != nil {
		return args.Get(0).(repository.Store), args.Error(1)
	}

	return nil, args.Error(1)
}

func (s *Store) Commit() error {
	args := s.Called()

	return args.Error(0)
}

func (s *Store) Rollback() error {
	args := s.Called()

	return args.Error(0)
}

func (s *Store) Item() repository.Item {
	return s.ItemMock
}

func (s *Store) User() repository.User {
	return s.UserMock
}
