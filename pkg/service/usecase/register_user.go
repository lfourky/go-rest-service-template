package usecase

import (
	"fmt"

	"github.com/lfourky/go-rest-service-template/pkg/model"
	"github.com/lfourky/go-rest-service-template/pkg/repository"
	"github.com/lfourky/go-rest-service-template/pkg/service"
	"github.com/pkg/errors"
)

type RegisterUser struct {
	store      repository.Store
	mailSender service.MailSender
}

func (u *RegisterUser) RegisterUser(email, name string) error {
	tx, txErr := u.store.BeginTransaction()
	if txErr != nil {
		return txErr
	}

	var err error
	defer func() {
		rollbackOnError(tx, err)
	}()

	var user *model.User
	user, err = tx.Users().FindByEmail(email)
	if err != nil {
		return err
	}

	if user != nil {
		err = errors.New("user already exists")
		return err
	}

	if err = tx.Users().Create(
		&model.User{
			Email: email,
			Name:  name,
		},
	); err != nil {
		return err
	}

	subject := "Successful registration"
	body := fmt.Sprintf("Thanks for registering, %s", name)
	if err = u.mailSender.SendMail(email, subject, body); err != nil {
		return err
	}

	if err = tx.Commit(); err != nil {
		return err
	}

	return nil
}
