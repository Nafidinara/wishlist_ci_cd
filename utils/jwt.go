package utils

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt"

	"go-wishlist-api/entities"
)

type JWTClaims struct {
	Id    int
	Email string
	jwt.StandardClaims
}

func GenerateToken(user *entities.User) (string, error) {

	secretToken := []byte(os.Getenv("SECRET_TOKEN"))
	claims := JWTClaims{
		Id:    user.ID,
		Email: user.Email,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(1 * time.Hour).Unix(),
			NotBefore: time.Now().Unix(),
			IssuedAt:  time.Now().Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedString, err := token.SignedString(secretToken)

	if err != nil {
		return "", err
	}
	return signedString, nil
}
