package services

import (
	"core-api/internal/domain/entities"
	"core-api/internal/infrastructure/repository"
	"fmt"
)

type BuildingService interface {
	CreateBuilding(building *entities.Building) error
	GetAllBuildings(accountID uint) ([]entities.Building, error)
	UpdateBuilding(buildingID uint, updatedBuilding entities.Building, accountID uint) error
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

// UpdateBuilding updates an existing building record.
func (s *buildingService) UpdateBuilding(buildingID uint, updatedBuilding entities.Building, accountID uint) error {
	// Optional: verify that the building belongs to the account.
	buildings, err := s.buildingRepo.FindAllByAccountID(accountID)
	if err != nil {
		return err
	}
	var exists bool
	for _, b := range buildings {
		if b.ID == buildingID {
			exists = true
			break
		}
	}
	if !exists {
		return fmt.Errorf("building not found for this account")
	}
	return s.buildingRepo.UpdateBuilding(buildingID, updatedBuilding)
}
