package usecase_test

import (
	"fmt"
	"testing"

	"github.com/lfourky/go-transaction-management/pkg/model"
	repomocks "github.com/lfourky/go-transaction-management/pkg/repository/mocks"
	"github.com/stretchr/testify/assert"
)

func TestRegisterUser(t *testing.T) {
	assert := assert.New(t)

	t.Run("user gets registered", func(t *testing.T) {
		u, store, _, userRepo, mailSender := setupUsecase()

		txStore := &repomocks.Store{}

		userName := "user-name"
		userEmail := "user-email"

		subject := "Successful registration"
		body := fmt.Sprintf("Thanks for registering, %s", userName)

		store.On("BeginTransaction").Return(txStore, nil)
		txStore.On("Users").Return(userRepo)
		userRepo.On("FindByEmail", userEmail).Return(nil, nil)
		userRepo.On("Create", &model.User{
			Name:  userName,
			Email: userEmail,
		}).Return(nil)
		mailSender.On("SendMail", userEmail, subject, body).Return(nil)
		txStore.On("Commit").Return(nil)

		err := u.RegisterUser(userEmail, userName)
		assert.NoError(err)

		txStore.AssertExpectations(t)
	})

	t.Run("user doesn't get registered due to tx error", func(t *testing.T) {
		u, store, _, userRepo, _ := setupUsecase()

		txStore := &repomocks.Store{}

		userName := "user-name"
		userEmail := "user-email"

		store.On("BeginTransaction").Return(txStore, nil)
		txStore.On("Users").Return(userRepo)
		userRepo.On("FindByEmail", userEmail).Return(nil, unexpectedError)
		txStore.On("Rollback").Return(nil)

		err := u.RegisterUser(userEmail, userName)
		assert.EqualError(err, unexpectedError.Error())

		txStore.AssertExpectations(t)
	})
}
