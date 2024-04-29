package wishlist

import (
	"time"

	"gorm.io/gorm"

	"go-wishlist-api/entities"
)

type Wishlist struct {
	ID         int `gorm:"primaryKey"`
	Title      string
	IsAchieved bool
	CreatedAt  time.Time `gorm:"autoCreateTime"`
	UpdatedAt  time.Time `gorm:"autoUpdateTime"`
	DeletedAt  gorm.DeletedAt
}

func FromUseCase(wishlist *entities.Wishlist) *Wishlist {
	return &Wishlist{
		ID:         wishlist.ID,
		Title:      wishlist.Title,
		IsAchieved: wishlist.IsAchieved,
		CreatedAt:  wishlist.CreatedAt,
		UpdatedAt:  wishlist.UpdatedAt,
	}
}

func FromUseCases(wishlists []entities.Wishlist) []Wishlist {
	var wishlistDb []Wishlist
	for _, wishlist := range wishlists {
		wishlistDb = append(wishlistDb, *FromUseCase(&wishlist))
	}
	return wishlistDb
}

func (w *Wishlist) ToUseCase() *entities.Wishlist {
	return &entities.Wishlist{
		ID:         w.ID,
		Title:      w.Title,
		IsAchieved: w.IsAchieved,
		CreatedAt:  w.CreatedAt,
		UpdatedAt:  w.UpdatedAt,
	}
}
