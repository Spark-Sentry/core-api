package repository

import (
	"core-api/internal/domain/entities"
	"errors"
	"fmt"
	"gorm.io/gorm"
)

type AccountRepository struct {
	db *gorm.DB
}

func NewAccountRepository(db *gorm.DB) *AccountRepository {
	return &AccountRepository{db: db}
}

// Create inserts a new account into the database.
func (repo *AccountRepository) Create(account *entities.Account) error {
	if err := repo.db.Create(account).Error; err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}

// FindByID finds an account by its ID.
func (repo *AccountRepository) FindByID(id uint) (*entities.Account, error) {
	var account entities.Account
	result := repo.db.First(&account, id)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		// Specific handling for not found error
		return nil, nil
	} else if result.Error != nil {
		// Handle other possible errors
		return nil, result.Error
	}
	return &account, nil
}

// Update modifies an existing account in the database.
func (repo *AccountRepository) Update(account *entities.Account) error {
	result := repo.db.Save(account)
	if result.Error != nil {
		// Handle errors, e.g., constraint violations
		return result.Error
	}
	return nil
}

// Delete removes an account from the database by ID.
func (repo *AccountRepository) Delete(id uint) error {
	result := repo.db.Delete(&entities.Account{}, id)
	if result.Error != nil {
		// Handle errors, possibly log them
		return result.Error
	}
	return nil
}

// List returns all accounts from the database.
func (repo *AccountRepository) List() ([]entities.Account, error) {
	var accounts []entities.Account
	result := repo.db.Find(&accounts)
	if result.Error != nil {
		// Handle potential errors, such as database connectivity issues
		return nil, result.Error
	}
	return accounts, nil
}

// FindByEmail finds an account by its contact email.
func (repo *AccountRepository) FindByEmail(email string) (*entities.Account, error) {
	var account entities.Account
	result := repo.db.Where("contact_email = ?", email).First(&account)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		// Return nil if no account is found to indicate absence without error
		return nil, nil
	} else if result.Error != nil {
		// Handle other errors, such as database issues
		return nil, result.Error
	}
	return &account, nil
}
