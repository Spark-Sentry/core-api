package services

import (
	"core-api/internal/infrastructure/repository"
	"errors"
	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
	"os"
	"time"
)

// Claims Définition de la structure pour les claims JWT
type Claims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

// AuthService struct pour la gestion de l'authentification
type AuthService struct {
	userRepo repository.UserRepository
}

// NewAuthService Fonction pour créer une nouvelle instance de AuthService
func NewAuthService(userRepo repository.UserRepository) *AuthService {
	return &AuthService{userRepo: userRepo}
}

// Authenticate Méthode pour authentifier un utilisateur et générer un JWT
func (s *AuthService) Authenticate(username, password string) (string, error) {
	user, err := s.userRepo.FindByUsername(username)
	if err != nil {
		return "", errors.New("user not found")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return "", errors.New("invalid password")
	}

	expirationTime := time.Now().Add(24 * time.Hour)
	claims := &Claims{
		Username: user.Name,
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
