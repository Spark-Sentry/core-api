package repository

import (
	"core-api/internal/domain/entities"
	"errors"
	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

// NewUserRepository creates a new instance of UserRepository.
func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{db: db}
}

// FindByEmail searches for a user by their email.
// It returns a pointer to a User entity if found, nil if not found, and an error for any other db related errors.
func (repo *UserRepository) FindByEmail(email string) (*entities.User, error) {
	var user entities.User
	if err := repo.db.Preload("Account").Where("email = ?", email).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil // Returns nil if the user is not found, without an error
		}
		return nil, err // Returns an error for other types of database errors
	}
	return &user, nil
}

// Save inserts a new user into the database.
// It takes a pointer to a User entity and attempts to create a record in the database.
// Returns an error if the operation fails.
func (repo *UserRepository) Save(user *entities.User) error {
	if err := repo.db.Create(user).Error; err != nil {
		return err // Returns an error if the saving operation fails
	}
	return nil
}

func (repo *UserRepository) AssociateUserToAccount(userID uint, accountID uint) error {
	user := &entities.User{}
	if err := repo.db.First(user, userID).Error; err != nil {
		return err
	}
	user.AccountID = &accountID
	return repo.db.Save(user).Error
}
