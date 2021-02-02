package mock

import (
	"github.com/lfourky/go-rest-service-template/pkg/model/domain"
	"github.com/stretchr/testify/mock"
)

type Item struct {
	mock.Mock
}

func (m *Item) Create(item *domain.Item) error {
	args := m.Called(item)

	return args.Error(0)
}

func (m *Item) FindAll() ([]*domain.Item, error) {
	args := m.Called()

	if args.Get(0) != nil {
		return args.Get(0).([]*domain.Item), args.Error(1)
	}

	return nil, args.Error(1)
}
