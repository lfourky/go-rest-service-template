package usecase_test

import (
	"errors"

	repomock "github.com/lfourky/go-rest-service-template/pkg/repository/mock"
	"github.com/lfourky/go-rest-service-template/pkg/service/log"
	servicemock "github.com/lfourky/go-rest-service-template/pkg/service/mock"
	"github.com/lfourky/go-rest-service-template/pkg/usecase"
)

var (
	errUnexpected = errors.New("unexpected error")
)

type usecaseSuite struct {
	store      *repomock.Store
	uc         *usecase.UseCase
	mailSender *servicemock.MailSender
}

func setupUsecase() *usecaseSuite {
	store := repomock.New()
	mailSender := &servicemock.MailSender{}

	return &usecaseSuite{
		uc:         usecase.New(store, mailSender, log.TestLogger()),
		store:      store,
		mailSender: mailSender,
	}
}
