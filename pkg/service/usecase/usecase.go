package usecase

import (
	"github.com/lfourky/go-rest-service-template/pkg/model"
	"github.com/lfourky/go-rest-service-template/pkg/repository"
	"github.com/lfourky/go-rest-service-template/pkg/service"
)

type Usecase struct {
	findUser     *FindUser
	registerUser *RegisterUser

	createItem *CreateItem
}

func New(store repository.Store, mailSender service.MailSender) *Usecase {
	return &Usecase{
		findUser:     &FindUser{store: store},
		registerUser: &RegisterUser{store: store, mailSender: mailSender},
		createItem:   &CreateItem{store: store},
	}
}

func (u *Usecase) FindUserByID(id string) (*model.User, error) {
	return u.findUser.FindUserByID(id)
}

func (u *Usecase) FindUserByEmail(email string) (*model.User, error) {
	return u.findUser.FindUserByEmail(email)
}

func (u *Usecase) RegisterUser(email, name string) error {
	return u.registerUser.RegisterUser(email, name)
}

func (u *Usecase) CreateItem(name string) (*model.Item, error) {
	return u.createItem.CreateItem(name)
}
