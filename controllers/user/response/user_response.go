package response

import "go-wishlist-api/entities"

type RegisterResponse struct {
	Name string `json:"name"`
	Email string `json:"email"`
}

type LoginResponse struct {
	Token string `json:"token"`
}

func FromUseCaseLogin(user *entities.Token) *LoginResponse {
	return &LoginResponse{
		Token: user.Token,
	}
}

func FromUseCaseRegister(user *entities.User) *RegisterResponse {
	return &RegisterResponse{
		Name: user.Name,
		Email: user.Email,
	}
}
