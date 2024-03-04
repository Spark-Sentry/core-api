package utils

import (
	"core-api/internal/domain/entities"
	"fmt"
	"github.com/golang-jwt/jwt"
	"os"
	"time"
)

// Claims struct of claims
type Claims struct {
	Email       string `json:"email"`
	Role        string `json:"role"`
	AccountName string `json:"accountName"`
	jwt.StandardClaims
}

func CreateJwt(user *entities.User) (string, error) {
	expirationTime := time.Now().Add(24 * time.Hour)
	fmt.Println(user.Account)
	claims := &Claims{
		Email:       user.Email,
		Role:        user.Role,
		AccountName: user.Account.Name,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(os.Getenv("JWT_SECRET_KEY")))
	if err != nil {
		return "", err
	}
	return tokenString, nil
}
