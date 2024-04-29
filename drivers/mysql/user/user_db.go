package user

import (
	"time"

	"gorm.io/gorm"

	"go-wishlist-api/entities"
)

type User struct {
	ID        int `gorm:"primaryKey"`
	Username  string
	Email     string
	Password  string
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
	DeletedAt gorm.DeletedAt
}

func FromUseCase(user *entities.User) *User {
	return &User{
		ID:        user.ID,
		Username:  user.Name,
		Email:     user.Email,
		Password:  user.Password,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}
}

func (user *User) ToUseCase() *entities.User {
	return &entities.User{
		ID:        user.ID,
		Name:      user.Username,
		Email:     user.Email,
		Password:  user.Password,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}
}
