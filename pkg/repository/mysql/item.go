package mysql

import (
	"errors"

	"github.com/lfourky/db-access/pkg/model"
)

type Item struct {
	*Repository
}

// This method demonstrates how gorm.DB is used in this architecture.
func (i *Item) FindItemByID(id string) (*model.Item, error) {
	var item *model.Item
	err := i.db.Find(&item).Where("id = ?", id).Error
	return item, err
}

func (i *Item) FindItemByName(name string) (*model.Item, error) {
	return nil, errors.New("not implemented")
}
