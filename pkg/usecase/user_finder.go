package usecase

import (
	"github.com/lfourky/go-rest-service-template/pkg/model/domain"
	"github.com/lfourky/go-rest-service-template/pkg/model/dto"
	"github.com/lfourky/go-rest-service-template/pkg/repository"
	"github.com/lfourky/go-rest-service-template/pkg/service/log"
)

type userFinder struct {
	store repository.Store

	logger *log.Logger
}

func newUserFinder(store repository.Store, logger *log.Logger) *userFinder {
	return &userFinder{
		store:  store,
		logger: logger,
	}
}

func (uc *userFinder) FindAllUsers() ([]*dto.User, error) {
	users, err := uc.store.User().FindAll()
	if err != nil {
		uc.logger.WithError(err).Warn("unable to find users")

		return nil, ErrDatabaseInternal
	}

	return uc.createResponse(users), nil
}

func (uc *userFinder) createResponse(users []*domain.User) []*dto.User {
	dtos := make([]*dto.User, 0, len(users))

	for _, u := range users {
		items := make([]*dto.Item, 0, len(u.Items))
		for _, i := range u.Items {
			items = append(items, &dto.Item{
				ID:   i.ID.String(),
				Name: i.Name,
			})
		}

		dtos = append(dtos, &dto.User{
			ID:    u.ID.String(),
			Name:  u.Name,
			Email: u.Email,
			Items: items,
		})
	}

	return dtos
}
