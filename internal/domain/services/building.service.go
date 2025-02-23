package services

import (
	"core-api/internal/domain/entities"
	"core-api/internal/infrastructure/repository"
	"fmt"
)

// BuildingService defines building-related services.
type BuildingService interface {
	CreateBuilding(building *entities.Building) error
	GetAllBuildings(accountID uint) ([]entities.Building, error)
}

type buildingService struct {
	buildingRepo repository.BuildingRepository
}

// NewBuildingService creates a new instance of BuildingService.
func NewBuildingService(buildingRepo repository.BuildingRepository) BuildingService {
	return &buildingService{
		buildingRepo: buildingRepo,
	}
}

// CreateBuilding creates a new Building record.
// CHANGES:
// - Validates that Name, Address, and CategoryID are provided.
// - For CategoryID, it checks that the pointer is not nil and the value is not 0.
func (s *buildingService) CreateBuilding(building *entities.Building) error {
	if building.Name == "" || building.Address == "" || building.CategoryID == nil || *building.CategoryID == 0 {
		return fmt.Errorf("invalid building data")
	}
	return s.buildingRepo.CreateBuilding(building)
}

// GetAllBuildings retrieves all buildings associated with the given account ID.
func (s *buildingService) GetAllBuildings(accountID uint) ([]entities.Building, error) {
	return s.buildingRepo.FindAllByAccountID(accountID)
}
