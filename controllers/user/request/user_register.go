package request

import "go-wishlist-api/entities"

type UserRegister struct {
	Name string `json:"name"`
	Email string `json:"email"`
	Password string `json:"password"`
}

func (r *UserRegister) ToEntities() *entities.User {
	return &entities.User{
		Name:       r.Name,
		Email:      r.Email,
		Password:   r.Password,
	}
}