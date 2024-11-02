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
	userRepo    repository.UserRepository
	accountRepo repository.AccountRepository
}

// NewAuthService creates a new AuthService with UserRepository and AccountRepository.
func NewAuthService(userRepo repository.UserRepository, accountRepo repository.AccountRepository) *AuthService {
	return &AuthService{
		userRepo:    userRepo,
		accountRepo: accountRepo,
	}
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

// RegisterUserAndAccount creates a new account and associates the user with it.
func (s *AuthService) RegisterUserAndAccount(user *entities.User, account *entities.Account) error {
	// Hash the user password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user.Password = string(hashedPassword)

	// Create the account in the database
	if err := s.accountRepo.Create(account); err != nil {
		return err
	}

	// Link the user to the newly created account
	user.AccountID = &account.ID

	// Save the user with the associated account
	return s.userRepo.Save(user)
}
