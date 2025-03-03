package services

import (
	"core-api/internal/app/dto"
	"core-api/internal/domain/entities"
	"core-api/internal/infrastructure/repository"
	"fmt"
)

// ProjectService defines operations related to projects.
type ProjectService interface {
	CreateProject(req dto.CreateProjectRequest) (*entities.Project, error)
	ListAllProjects() ([]entities.Project, error)
	GetProjectByID(id uint) (*entities.Project, error)
	UpdateProjectByID(id uint, req dto.UpdateProjectRequest) error
}

type projectService struct {
	projectRepo repository.ProjectRepository
	targetRepo  repository.TargetRepository
}

// NewProjectService creates a new instance of ProjectService.
func NewProjectService(projectRepo repository.ProjectRepository, targetRepo repository.TargetRepository) ProjectService {
	return &projectService{
		projectRepo: projectRepo,
		targetRepo:  targetRepo,
	}
}

// CreateProject creates a new project along with its associated targets.
func (s *projectService) CreateProject(req dto.CreateProjectRequest) (*entities.Project, error) {
	project := &entities.Project{
		BuildingID:         req.BuildingID,
		Name:               req.Name,
		ContractorID:       req.ContractorID,
		ImplementationDate: req.ImplementationDate,
		EfficiencyCost:     req.EfficiencyCost,
		MaintenanceCost:    req.MaintenanceCost,
	}

	if err := s.projectRepo.CreateProject(project); err != nil {
		return nil, fmt.Errorf("failed to create project: %w", err)
	}

	// Create each target and associate it with the project.
	for _, t := range req.Targets {
		target := entities.Target{
			Name:      t.Name,
			Value:     t.Value,
			Type:      t.Type,
			ProjectID: &project.ID,
		}
		if err := s.targetRepo.CreateTarget(&target); err != nil {
			return nil, fmt.Errorf("failed to create target: %w", err)
		}
		project.Targets = append(project.Targets, target)
	}

	return project, nil
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
