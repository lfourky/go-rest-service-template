package usecase

import (
	"fmt"

	"github.com/lfourky/go-rest-service-template/pkg/model/domain"
	"github.com/lfourky/go-rest-service-template/pkg/model/dto"
	"github.com/lfourky/go-rest-service-template/pkg/repository"
	"github.com/lfourky/go-rest-service-template/pkg/service"
	"github.com/lfourky/go-rest-service-template/pkg/service/log"
)

type registerUser struct {
	store      repository.Store
	mailSender service.MailSender

	logger *log.Logger
}

func newRegisterUser(
	store repository.Store,
	mailSender service.MailSender,
	logger *log.Logger,
) *registerUser {
	return &registerUser{
		store:      store,
		mailSender: mailSender,
		logger:     logger,
	}
}

func (uc *registerUser) RegisterUser(request *dto.RegisterUserRequest) (*dto.RegisterUserResponse, error) {
	tx, err := beginTx(uc.store, uc.logger)
	if err != nil {
		return nil, err
	}

	defer rollback(tx, uc.logger)

	user, err := tx.User().FindByEmail(request.Email)
	if err != nil {
		uc.logger.WithError(err).Warn("unable to find user")

		return nil, ErrDatabaseInternal
	}

	if user != nil {
		return nil, ErrDatabaseUserAlreadyExists
	}

	user = &domain.User{
		Email: request.Email,
		Name:  request.Name,
	}
	if err := tx.User().Create(user); err != nil {
		uc.logger.WithError(err).Warn("unable to create user")

		return nil, ErrDatabaseInternal
	}

	var (
		subject = "Successful registration"
		body    = fmt.Sprintf("Thanks for registering, %s", user.Name)
	)

	if err = uc.mailSender.SendMail(user.Email, subject, body); err != nil {
		uc.logger.WithError(err).Warn("unable to send mail")

		return nil, ErrDomainMailSendingFailed
	}

	if err := commit(tx, uc.logger); err != nil {
		return nil, err
	}

	return &dto.RegisterUserResponse{
		User: dto.User{
			ID:    user.ID.String(),
			Name:  user.Name,
			Email: user.Email,
		},
	}, nil
}
