package services

import (
	"core-api/internal/domain/entities"
	"core-api/internal/infrastructure/repository"
)

type AccountService struct {
	userRepo    repository.UserRepository
	accountRepo repository.AccountRepository
}

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
