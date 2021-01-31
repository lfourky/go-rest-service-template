package usecase

import (
	"github.com/lfourky/go-rest-service-template/pkg/model/domain"
	"github.com/lfourky/go-rest-service-template/pkg/model/dto"
	"github.com/lfourky/go-rest-service-template/pkg/repository"
	"github.com/lfourky/go-rest-service-template/pkg/service/log"
)

type itemCreation struct {
	store repository.Store

	logger *log.Logger
}

func newItemCreation(store repository.Store, logger *log.Logger) *itemCreation {
	return &itemCreation{
		store:  store,
		logger: logger,
	}
}

func (uc *itemCreation) CreateItem(request *dto.CreateItemRequest) (*dto.CreateItemResponse, error) {
	user, err := findUserByID(request.UserID, uc.store, uc.logger)
	if err != nil {
		return nil, err
	}

	item := &domain.Item{
		Name:   request.Name,
		UserID: user.ID,
	}

	if err := uc.store.Item().Create(item); err != nil {
		uc.logger.WithError(err).Warn("unable to create item")

		return nil, ErrDatabaseItemCreationFailed
	}

	return &dto.CreateItemResponse{
		Item: dto.Item{
			ID:   item.ID.String(),
			Name: item.Name,
		},
	}, nil
}
