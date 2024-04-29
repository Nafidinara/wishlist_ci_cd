package entities

import "time"

type Wishlist struct {
	ID         int
	Title      string
	IsAchieved bool
	CreatedAt  time.Time
	UpdatedAt  time.Time
	DeletedAt  time.Time
}

type IWishlistRepository interface {
	GetAll() ([]Wishlist, error)
	Create(wishlist *Wishlist) error
}

type IWishlistUseCase interface {
	GetAll() ([]Wishlist, error)
	Create(wishlist *Wishlist) (*Wishlist, error)
}
