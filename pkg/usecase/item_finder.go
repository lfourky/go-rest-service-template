package usecase

import (
	"github.com/lfourky/go-rest-service-template/pkg/model/domain"
	"github.com/lfourky/go-rest-service-template/pkg/model/dto"
	"github.com/lfourky/go-rest-service-template/pkg/repository"
	"github.com/lfourky/go-rest-service-template/pkg/service/log"
)

type itemFinder struct {
	store repository.Store

	logger *log.Logger
}

func newItemFinder(store repository.Store, logger *log.Logger) *itemFinder {
	return &itemFinder{
		store:  store,
		logger: logger,
	}
}

func (uc *itemFinder) FindAllItems() ([]*dto.Item, error) {
	items, err := uc.store.Item().FindAll()
	if err != nil {
		uc.logger.WithError(err).Warn("unable to find items")

		return nil, ErrDatabaseInternal
	}

	return uc.createResponse(items), nil
}

func (uc *itemFinder) createResponse(items []*domain.Item) []*dto.Item {
	dtos := make([]*dto.Item, 0, len(items))

	for _, i := range items {
		dtos = append(dtos, &dto.Item{
			ID:     i.ID.String(),
			Name:   i.Name,
			UserID: i.UserID.String(),
		})
	}

	return dtos
}
