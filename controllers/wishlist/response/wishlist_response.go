package response

import "go-wishlist-api/entities"

type WishlistResponse struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
	IsAchieved bool `json:"is_achieved"`
}

func FromUseCase(wishlist *entities.Wishlist) *WishlistResponse {
	return &WishlistResponse{
		ID:          wishlist.ID,
		Title:       wishlist.Title,
		IsAchieved: wishlist.IsAchieved,
	}
}

func FromUseCases(wishlists []entities.Wishlist) []WishlistResponse {
	var wishlistResponses []WishlistResponse
	for _, wishlist := range wishlists {
		wishlistResponses = append(wishlistResponses, *FromUseCase(&wishlist))
	}
	return wishlistResponses
}