package repositories

import (
	"gorm.io/gorm"

	db_user "go-wishlist-api/drivers/mysql/user"
	"go-wishlist-api/entities"
)

type UserRepo struct {
	DB *gorm.DB
}

func NewUserRepo(db *gorm.DB) *UserRepo {
	return &UserRepo{DB: db}
}

func (r *UserRepo) FindByEmail(email string) (*entities.User, error) {
	userDb := entities.User{}

	if err := r.DB.Where("email = ?", email).First(&userDb).Error; err != nil {
		return &userDb, err
	}
	return &userDb, nil
}

func (r *UserRepo) Create(user *entities.User) (*entities.User, error) {
	userDb := db_user.FromUseCase(user)

	if err := r.DB.Create(&userDb).Error; err != nil {
		return nil, err
	}

	return userDb.ToUseCase(), nil

}
