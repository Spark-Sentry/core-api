package repository

import (
	"core-api/internal/domain/entities"
	"gorm.io/gorm"
)

type AccountRepository interface {
	CreateAccount(account *entities.Account) error
}

type GormAccountRepository struct {
	db *gorm.DB
}

func (repo *GormAccountRepository) CreateAccount(account *entities.Account) error {
	return repo.db.Create(account).Error
}
