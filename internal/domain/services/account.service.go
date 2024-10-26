package services

import (
	"core-api/internal/domain/entities"
	"core-api/internal/infrastructure/repository"
	"errors"
)

type AccountService struct {
	userRepo    repository.UserRepository
	accountRepo repository.AccountRepository
}

// ErrAccountNotFound is returned when an account cannot be found in the repository.
var ErrAccountNotFound = errors.New("account not found")

func NewAccountService(userRepo repository.UserRepository, accountRepo repository.AccountRepository) *AccountService {
	return &AccountService{
		userRepo:    userRepo,
		accountRepo: accountRepo,
	}
}

func (s *AccountService) AssociateUserToAccount(userID, accountID uint) error {
	return s.userRepo.AssociateUserToAccount(userID, accountID)
}

func (s *AccountService) CreateAccount(account *entities.Account) error {
	return s.accountRepo.Create(account)
}

// GetAllAccounts retrieves all accounts from the repository.
func (s *AccountService) GetAllAccounts() ([]entities.Account, error) {
	accounts, err := s.accountRepo.List()
	if err != nil {
		return nil, err
	}
	return accounts, nil
}

// GetAccountByID retrieves an account by its ID from the repository.
func (s *AccountService) GetAccountByID(id uint) (*entities.Account, error) {
	account, err := s.accountRepo.FindByID(id)
	if err != nil {
		return nil, err
	}
	return account, nil
}
