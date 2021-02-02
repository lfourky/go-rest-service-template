package mysql

import (
	"github.com/lfourky/go-rest-service-template/pkg/model/domain"
	"gorm.io/gorm"
)

type Item struct {
	db *gorm.DB
}

func (m *Item) Create(item *domain.Item) error {
	return m.db.Create(item).Error
}

func (m *Item) Delete(item *domain.Item) error {
	return m.db.Delete(item).Error
}

func (m *Item) FindAll() ([]*domain.Item, error) {
	var items []*domain.Item

	return items, m.db.Find(&items).Error
}
