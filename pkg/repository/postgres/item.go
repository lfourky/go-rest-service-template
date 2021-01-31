package postgres

import (
	"github.com/lfourky/go-rest-service-template/pkg/model/domain"
	"gorm.io/gorm"
)

type Item struct {
	db *gorm.DB
}

func (p *Item) Create(item *domain.Item) error {
	return p.db.Create(item).Error
}

func (p *Item) FindAll() ([]*domain.Item, error) {
	var items []*domain.Item
	return items, p.db.Find(&items).Error
}
