package services

import (
	"core-api/internal/app/dto"
	"core-api/internal/domain/entities"
	"core-api/internal/infrastructure/repository"
	"fmt"
)

// ProjectService defines operations related to projects.
type ProjectService interface {
	CreateProject(project *entities.Project) error
	ListAllProjects() ([]entities.Project, error)
	GetProjectByID(id uint) (*entities.Project, error)
	UpdateProjectByID(id uint, req dto.UpdateProjectRequest) error
}

type projectService struct {
	projectRepo repository.ProjectRepository
}

// NewProjectService creates a new instance of ProjectService.
func NewProjectService(projectRepo repository.ProjectRepository) ProjectService {
	return &projectService{
		projectRepo: projectRepo,
	}
}

// CreateProject validates and creates a new project.
func (s *projectService) CreateProject(project *entities.Project) error {
	if project.Name == "" || project.BuildingID == 0 || project.ContractorID == 0 || project.ImplementationDate == "" {
		return fmt.Errorf("invalid project data")
	}
	return s.projectRepo.CreateProject(project)
}

// ListAllProjects retrieves all projects.
func (s *projectService) ListAllProjects() ([]entities.Project, error) {
	return s.projectRepo.ListAllProjects()
}

// GetProjectByID retrieves a project by its ID.
func (s *projectService) GetProjectByID(id uint) (*entities.Project, error) {
	return s.projectRepo.FindProjectByID(id)
}

// UpdateProjectByID directly updates a project by its ID using GORM.
func (s *projectService) UpdateProjectByID(id uint, req dto.UpdateProjectRequest) error {
	// Build a map with the fields to update.
	updateData := map[string]interface{}{
		"name":                req.Name,
		"contractor_id":       req.ContractorID,
		"implementation_date": req.ImplementationDate,
		"efficiency_cost":     req.EfficiencyCost,
		"maintenance_cost":    req.MaintenanceCost,
	}
	if err := s.projectRepo.UpdateProjectByID(id, updateData); err != nil {
		return fmt.Errorf("failed to update project: %w", err)
	}
	return nil
}
