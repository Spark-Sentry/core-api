package services

import (
	"core-api/internal/domain/entities"
	"core-api/internal/infrastructure/repository"
)

// CategoryService defines operations related to categories.
type CategoryService interface {
	ListCategories() ([]entities.Category, error)
}

type categoryService struct {
	repo repository.CategoryRepository
}

// NewCategoryService creates a new instance of CategoryService.
func NewCategoryService(repo repository.CategoryRepository) CategoryService {
	return &categoryService{
		repo: repo,
	}
}

// ListCategories retrieves all categories from the repository.
func (s *categoryService) ListCategories() ([]entities.Category, error) {
	return s.repo.ListAllCategories()
}
