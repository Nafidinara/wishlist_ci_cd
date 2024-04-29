package request

import "go-wishlist-api/entities"

type WishlistCreate struct {
	ID int `json:"id"`
	Title string `json:"title"`
	IsAchieved bool `json:"is_achieved"`
}

func (r *WishlistCreate) ToEntities() *entities.Wishlist {
	return &entities.Wishlist{
		ID:         r.ID,
		Title:      r.Title,
		IsAchieved: r.IsAchieved,
	}	
}