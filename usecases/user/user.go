package usecases

import (
	"go-wishlist-api/constant"
	"go-wishlist-api/entities"
	"go-wishlist-api/utils"
)

type UserUseCase struct {
	repository entities.IUserRepository
}

func NewUserUseCase(repository entities.IUserRepository) *UserUseCase {
	return &UserUseCase{
		repository: repository,
	}
}

func (u *UserUseCase) Register(userReq *entities.User) (*entities.User, error) {
	if userReq.Name == "" || userReq.Email == "" || userReq.Password == "" {
		return nil, constant.ErrEmptyInput
	}

	existUser, _ := u.repository.FindByEmail(userReq.Email)

	if existUser != nil {
		return nil, constant.ErrUserExist
	}

	pass, err := utils.HashPassword(userReq.Password)

	if err != nil {
		return nil, constant.ErrUserCreate
	}

	user := &entities.User{
		Name:     userReq.Name,
		Email:    userReq.Email,
		Password: pass,
	}

	newUser, err := u.repository.Create(user)

	if err != nil {
		return nil, constant.ErrUserCreate
	}

	return newUser, nil
}

func (u *UserUseCase) Login(userReq *entities.User) (*entities.Token, error) {
	if userReq.Email == "" || userReq.Password == "" {
		return nil, constant.ErrEmptyInput
	}

	existUser, _ := u.repository.FindByEmail(userReq.Email)

	if existUser == nil {
		return nil, constant.ErrUserNotExist
	}

	if err := utils.VerifyPassword(userReq.Password, existUser.Password); err != nil {
		return nil, constant.ErrUserWrongInput
	}

	token, err := utils.GenerateToken(existUser)

	if err != nil {
		return nil, constant.ErrInternalServer
	}

	return &entities.Token{Token: token}, nil
}
