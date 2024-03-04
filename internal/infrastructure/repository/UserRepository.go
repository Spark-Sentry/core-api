package repository

import (
	"core-api/internal/domain/entities"
	"gorm.io/gorm"
)

// UserRepository struct pour gérer les opérations de base de données des utilisateurs
type UserRepository struct {
	db *gorm.DB
}

// NewUserRepository Fonction pour créer une nouvelle instance de UserRepository
func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{db: db}
}

// FindByUsername Méthode pour trouver un utilisateur par son nom d'utilisateur
func (repo *UserRepository) FindByUsername(username string) (*entities.User, error) {
	var user entities.User
	if err := repo.db.Where("name = ?", username).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}
