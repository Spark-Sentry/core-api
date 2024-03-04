package services

import (
	"core-api/internal/domain/entities"
	"core-api/internal/infrastructure/repository"
)

type UserService struct {
	userRepo    repository.UserRepository
	accountRepo repository.AccountRepository
}

func NewUserService(userRepo repository.UserRepository, accountRepo repository.AccountRepository) *UserService {
	return &UserService{
		userRepo:    userRepo,
		accountRepo: accountRepo,
	}
}

func (s *UserService) GetUserDetails(userMail string) (*entities.User, error) {
	user, err := s.userRepo.FindByEmail(userMail)
	if err != nil {
		return nil, err
	}
	return user, nil
}
