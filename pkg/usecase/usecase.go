package usecase

import (
	"github.com/lfourky/go-rest-service-template/pkg/model/dto"
	"github.com/lfourky/go-rest-service-template/pkg/repository"
	"github.com/lfourky/go-rest-service-template/pkg/service"
	"github.com/lfourky/go-rest-service-template/pkg/service/log"
)

type User interface {
	FindAllUsers() ([]*dto.User, error)
	RegisterUser(request *dto.RegisterUserRequest) (*dto.RegisterUserResponse, error)
}

type Item interface {
	FindAllItems() ([]*dto.Item, error)
	CreateItem(request *dto.CreateItemRequest) (*dto.CreateItemResponse, error)
}

// UseCase is a helper structure that holds all the use cases embedded within it.
type UseCase struct {
	*userFinder
	*registerUser

	*itemFinder
	*itemCreation
}

func New(
	store repository.Store,
	mailSender service.MailSender,
	logger *log.Logger,
) *UseCase {
	return &UseCase{
		userFinder:   newUserFinder(store, logger),
		registerUser: newRegisterUser(store, mailSender, logger),
		itemCreation: newItemCreation(store, logger),
		itemFinder:   newItemFinder(store, logger),
	}
}
