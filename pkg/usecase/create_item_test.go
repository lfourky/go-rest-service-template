package usecase_test

import (
	"testing"

	"github.com/lfourky/go-rest-service-template/pkg/model/domain"
	"github.com/lfourky/go-rest-service-template/pkg/model/dto"
	"github.com/lfourky/go-rest-service-template/pkg/usecase"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestCreateItem(t *testing.T) {
	assert := assert.New(t)

	var (
		userUUID = domain.UUID("55a23280-6599-11eb-ae93-0242ac130002")

		request = &dto.CreateItemRequest{
			UserID: userUUID.String(),
			Name:   "item-name",
		}
	)

	type createItemSuite struct {
		*usecaseSuite

		user *domain.User
		item *domain.Item
	}

	setup := func() *createItemSuite {
		return &createItemSuite{
			usecaseSuite: setupUsecase(),

			user: &domain.User{
				PrimaryKey: domain.PrimaryKey{
					ID: userUUID,
				},
			},
			item: &domain.Item{
				Name:   "item-name",
				UserID: userUUID,
			},
		}
	}

	tests := []struct {
		name     string
		err      error
		request  *dto.CreateItemRequest
		response *dto.CreateItemResponse
		mocks    func(suite *createItemSuite)
	}{
		{
			name:    "item created successfully",
			request: request,
			response: &dto.CreateItemResponse{
				Item: dto.Item{
					Name: "item-name",
				},
			},
			mocks: func(suite *createItemSuite) {
				suite.store.UserMock.On("FindByID", suite.user.ID).Return(suite.user, nil)
				suite.store.ItemMock.On("Create", suite.item).Return(nil)
			},
		},
		{
			name:    "unable to create item due to unexpected error when creating item",
			request: request,
			err:     usecase.ErrDatabaseItemCreationFailed,
			mocks: func(suite *createItemSuite) {
				suite.store.UserMock.On("FindByID", suite.user.ID).Return(suite.user, nil)
				suite.store.ItemMock.On("Create", suite.item).Return(errUnexpected)
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			suite := setup()
			test.mocks(suite)
			resp, err := suite.uc.CreateItem(test.request)
			assert.Equal(test.err, err)
			assert.Equal(test.response, resp)
			mock.AssertExpectationsForObjects(t, suite.store, suite.store.ItemMock, suite.store.UserMock)
		})
	}
}
