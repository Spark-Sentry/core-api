package services

import (
	"core-api/internal/domain/entities"
	"core-api/internal/infrastructure/repository"
	"core-api/pkg/utils"
	"errors"
	"golang.org/x/crypto/bcrypt"
)

// AuthService struct for AuthService
type AuthService struct {
	userRepo repository.UserRepository
}

// NewAuthService method to create a new AuthService
func NewAuthService(userRepo repository.UserRepository) *AuthService {
	return &AuthService{userRepo: userRepo}
}

// Authenticate method to authenticate and create a new token
func (s *AuthService) Authenticate(email string, password string) (string, error) {
	user, err := s.userRepo.FindByEmail(email)
	if err != nil {
		return "", errors.New("user not found")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return "", errors.New("invalid password")
	}

	tokenString, err := utils.CreateJwt(user)

	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func (s *AuthService) RegisterUser(user *entities.User) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user.Password = string(hashedPassword)

	return s.userRepo.Save(user)
}
