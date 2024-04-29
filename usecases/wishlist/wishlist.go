package usecases

import (
	"go-wishlist-api/constant"
	"go-wishlist-api/entities"
)

type WishlistUseCase struct {
	repository entities.IWishlistRepository
}

func NewWishlistUseCase(repository entities.IWishlistRepository) *WishlistUseCase {
	return &WishlistUseCase{
		repository: repository,
	}
}

func (w *WishlistUseCase) GetAll() ([]entities.Wishlist, error) {
	return w.repository.GetAll()
}

func (w *WishlistUseCase) Create(wishlist *entities.Wishlist) (*entities.Wishlist, error) {
	if wishlist.Title == "" {
		return &entities.Wishlist{}, constant.ErrEmptyInput
	}

	err := w.repository.Create(wishlist)

	if err != nil {
		return &entities.Wishlist{}, constant.ErrInternalServer
	}

	return wishlist, nil
}
