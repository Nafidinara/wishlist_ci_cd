package repositories

import (
	"gorm.io/gorm"

	db_wishlist "go-wishlist-api/drivers/mysql/wishlist"
	"go-wishlist-api/entities"
)

type WishlistRepo struct {
	DB *gorm.DB
}

func NewWishlistRepo(db *gorm.DB) *WishlistRepo {
	return &WishlistRepo{DB: db}
}

func (r *WishlistRepo) GetAll() ([]entities.Wishlist, error) {
	wishlistsDb := []entities.Wishlist{}

	if err := r.DB.Find(&wishlistsDb).Error; err != nil {
		return nil, err
	}
	return wishlistsDb, nil
}

func (r *WishlistRepo) Create(wishlist *entities.Wishlist) (error) {
	wishlistDb := db_wishlist.FromUseCase(wishlist)

	if err := r.DB.Create(&wishlistDb).Error; err != nil {
		return err
	}

	return nil
}
