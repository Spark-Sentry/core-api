package repository

import (
	"core-api/internal/domain/entities"
	"gorm.io/gorm"
)

// ProjectRepository handles database operations for projects.
type ProjectRepository interface {
	CreateProject(project *entities.Project) error
	ListAllProjects() ([]entities.Project, error)
	FindProjectByID(id uint) (*entities.Project, error)
	UpdateProjectByID(id uint, updateData map[string]interface{}) error
}

type projectRepository struct {
	db *gorm.DB
}

// NewProjectRepository creates a new instance of ProjectRepository.
func NewProjectRepository(db *gorm.DB) ProjectRepository {
	return &projectRepository{db: db}
}

// CreateProject inserts a new project record into the database.
func (r *projectRepository) CreateProject(project *entities.Project) error {
	return r.db.Create(project).Error
}

// ListAllProjects retrieves all project records.
func (r *projectRepository) ListAllProjects() ([]entities.Project, error) {
	var projects []entities.Project
	err := r.db.Find(&projects).Error
	return projects, err
}

// FindProjectByID retrieves a project by its ID.
func (r *projectRepository) FindProjectByID(id uint) (*entities.Project, error) {
	var project entities.Project
	err := r.db.First(&project, id).Error
	if err != nil {
		return nil, err
	}
	return &project, nil
}

// UpdateProjectByID updates the project with the given ID using the provided updateData.
func (r *projectRepository) UpdateProjectByID(id uint, updateData map[string]interface{}) error {
	return r.db.Model(&entities.Project{}).Where("id = ?", id).Updates(updateData).Error
}
