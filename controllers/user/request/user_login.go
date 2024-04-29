package request

import "go-wishlist-api/entities"

type UserLogin struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (r *UserLogin) ToEntities() *entities.User {
	return &entities.User{
		Email:    r.Email,
		Password: r.Password,
	}
}
