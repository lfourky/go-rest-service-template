package mysql

import (
	"github.com/lfourky/go-transaction-management/pkg/model"
)

type Item struct {
	*Repository
}

func (i *Item) Create(item *model.Item) error {
	return i.db.Create(item).Error
}

func (i *Item) Delete(item *model.Item) error {
	return i.db.Delete(item).Error
}
