package usecases

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"go-wishlist-api/constant"
	"go-wishlist-api/drivers/mysql/mocks"
	"go-wishlist-api/entities"
)

func TestCreate(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		mockWishlist := &entities.Wishlist{
			ID:         1,
			Title:      "Wishlist Test",
			IsAchieved: true,
		}

		mockRepo := new(mocks.MockWishlistRepository)
		uc := NewWishlistUseCase(mockRepo)

		mockRepo.On("Create", mockWishlist).Return(mockWishlist, nil)
		wishlist, err := uc.Create(mockWishlist)

		assert.NoError(t, err)
		assert.NotNil(t, wishlist)
		assert.Equal(t, mockWishlist, wishlist)
		mockRepo.AssertExpectations(t)
	})

	t.Run("Failed", func(t *testing.T) {
		mockWishlist := &entities.Wishlist{
			Title:      "Wishlist Test",
			IsAchieved: true,
		}

		mockRepo := new(mocks.MockWishlistRepository)

		uc := NewWishlistUseCase(mockRepo)

		mockRepo.On("Create", mockWishlist).Return(nil, constant.ErrInternalServer)

		wishlist, err := uc.Create(mockWishlist)

		assert.Error(t, err)
		assert.Empty(t, wishlist)
		assert.EqualError(t, err, constant.ErrInternalServer.Error())
		mockRepo.AssertExpectations(t)
	})
}

func TestGetAll(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		mockWishlists := []entities.Wishlist{
			{ID: 1, Title: "Wishlist 1", IsAchieved: false},
			{ID: 2, Title: "Wishlist 2", IsAchieved: true},
		}
		mockRepo := new(mocks.MockWishlistRepository)
		uc := NewWishlistUseCase(mockRepo)
		mockRepo.On("GetAll").Return(mockWishlists, nil)
		wishlists, err := uc.GetAll()
		assert.NoError(t, err)
		assert.NotNil(t, wishlists)
		assert.Equal(t, len(mockWishlists), len(wishlists))
		mockRepo.AssertExpectations(t)
	})

	t.Run("Failed", func(t *testing.T) {
		mockRepo := new(mocks.MockWishlistRepository)
		uc := NewWishlistUseCase(mockRepo)
		mockRepo.On("GetAll").Return(nil, constant.ErrInternalServer)
		wishlists, err := uc.GetAll()
		assert.Error(t, err)
		assert.Empty(t, wishlists)
		assert.EqualError(t, err, constant.ErrInternalServer.Error())
		mockRepo.AssertExpectations(t)
	})
}