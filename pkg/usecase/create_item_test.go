package usecase_test

import (
	"testing"
)

func TestCreateItem(t *testing.T) {
	// assert := assert.New(t)

	// t.Run("item gets created", func(t *testing.T) {
	// 	u, store, itemRepo, _, _ := setupUsecase()

	// 	expectedItem := &domain.Item{
	// 		Name: "item_name",
	// 	}

	// 	store.On("Items").Return(itemRepo)
	// 	itemRepo.On("Create", &domain.Item{
	// 		Name: "item_name",
	// 	}).Return(nil)

	// 	item, err := u.CreateItem("item_name")
	// 	assert.NoError(err)
	// 	assert.Equal(expectedItem, item)
	// })

	// t.Run("item creation fails due to unexpected db error", func(t *testing.T) {
	// 	u, store, itemRepo, _, _ := setupUsecase()

	// 	store.On("Items").Return(itemRepo)
	// 	itemRepo.On("Create", &domain.Item{
	// 		Name: "item_name",
	// 	}).Return(unexpectedError)

	// 	item, err := u.CreateItem("item_name")
	// 	assert.EqualError(err, unexpectedError.Error())
	// 	assert.Nil(item)
	// })
}
