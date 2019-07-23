package service

import "github.com/lfourky/go-transaction-management/pkg/model"

type User interface {
	RegisterUser(email, name string) error
	FindUserByID(id string) (*model.User, error)
	FindUserByEmail(email string) (*model.User, error)
}

type Item interface {
	CreateItem(name string) (*model.Item, error)
}

type MailSender interface {
	SendMail(recipient, subject, body string) error
}
