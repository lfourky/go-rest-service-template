package mysql

import (
	"errors"

	"github.com/lfourky/go-transaction-management/pkg/model"
)

type Cart struct {
	*Repository
}

// This method demonstrates how gorm.DB is used in this architecture.
func (c *Cart) FindCartByID(id string) (*model.Cart, error) {
	var cart *model.Cart
	err := c.db.Find(&cart).Where("id = ?", id).Error
	return cart, err
}

func (c *Cart) FindCartByItem(i model.Item) (*model.Cart, error) {
	return nil, errors.New("not implemented")
}
