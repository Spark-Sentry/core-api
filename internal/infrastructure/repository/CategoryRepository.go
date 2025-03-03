package repository

import (
	"core-api/internal/domain/entities"
	"gorm.io/gorm"
)

// CategoryRepository handles database operations for categories.
type CategoryRepository interface {
	ListAllCategories() ([]entities.Category, error)
}

type categoryRepository struct {
	db *gorm.DB
}

// NewCategoryRepository creates a new instance of CategoryRepository.
func NewCategoryRepository(db *gorm.DB) CategoryRepository {
	return &categoryRepository{db: db}
}

// ListAllCategories retrieves all categories from the database.
func (r *categoryRepository) ListAllCategories() ([]entities.Category, error) {
	var categories []entities.Category
	err := r.db.Find(&categories).Error
	return categories, err
}
