package usecase_test

import (
	"errors"
)

var (
	// Used globally in tests - not to be altered.
	errUnexpected = errors.New("unexpected error")
)

// func setupUsecase() (*usecase.Usecase, *repomocks.Store, *repomocks.Item, *repomocks.User, *servicemocks.MailSender) {
// 	store := &repomocks.Store{}
// 	itemRepository := &repomocks.Item{}
// 	userRepository := &repomocks.User{}

// 	mailSender := &servicemocks.MailSender{}

// 	u := usecase.New(store, mailSender)
// 	return u, store, itemRepository, userRepository, mailSender
// }
