package usecase_test

import (
	"errors"

	repomocks "github.com/lfourky/go-transaction-management/pkg/repository/mocks"
	servicemocks "github.com/lfourky/go-transaction-management/pkg/service/mocks"
	"github.com/lfourky/go-transaction-management/pkg/service/usecase"
)

var (
	// Used globally in tests - not to be altered.
	unexpectedError = errors.New("unexpected error")
)

func setupUsecase() (*usecase.Usecase, *repomocks.Store, *repomocks.Item, *repomocks.User, *servicemocks.MailSender) {
	store := &repomocks.Store{}
	itemRepository := &repomocks.Item{}
	userRepository := &repomocks.User{}

	mailSender := &servicemocks.MailSender{}

	u := usecase.New(store, mailSender)
	return u, store, itemRepository, userRepository, mailSender
}
