package usecase

import (
	"github.com/lfourky/go-transaction-management/pkg/model"
	"github.com/lfourky/go-transaction-management/pkg/repository"
)

type CreateItem struct {
	store repository.Store
}

func (i *CreateItem) CreateItem(name string) (*model.Item, error) {
	item := &model.Item{
		Name: name,
	}

	if err := i.store.Items().Create(item); err != nil {
		return nil, err
	}

	return item, nil
}
