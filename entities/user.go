package entities

import "time"

// type Wishlist struct {
// 	ID         int
// 	Title      string
// 	IsAchieved bool
// 	CreatedAt  time.Time
// 	UpdatedAt  time.Time
// 	DeletedAt  time.Time
// }

// type RepositoryInterface interface {
// 	GetAll() ([]Wishlist, error)
// 	Create(wishlist *Wishlist) error
// }

// type UseCaseInterface interface {
// 	GetAll() ([]Wishlist, error)
// 	Create(wishlist *Wishlist) (Wishlist, error)
// }

type User struct {
	ID       int
	Name     string
	Email    string
	Password string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time
}

type Token struct {
	Token string
}

type IUserRepository interface {
	FindByEmail(email string) (*User, error)
	Create(user *User) (*User, error)
}

type IUserUseCase interface {
	Register(userReq *User) (*User, error)
	Login(userReq *User) (*Token, error)
}
