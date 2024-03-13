package services

import (
	"core-api/internal/app/dto"
	"core-api/internal/domain/entities"
	"core-api/internal/infrastructure/repository"
)

// BuildingService contains the business logic for Building entities.
type BuildingService struct {
	buildingRepo repository.BuildingRepository
}

// NewBuildingService creates a new instance of BuildingService.
func NewBuildingService(buildingRepo repository.BuildingRepository) *BuildingService {
	return &BuildingService{buildingRepo: buildingRepo}
}

// CreateBuilding creates a new Building with at least one Area from a DTO.
func (s *BuildingService) CreateBuilding(req *dto.CreateBuildingRequest, accountID uint) error {
	building := entities.Building{
		Name:      req.Name,
		Address:   req.Address,
		Group:     req.Group,
		AccountID: accountID,
	}

	building.Areas = []entities.Area{}

	for _, areaReq := range req.Areas {
		area := entities.Area{
			Name:        areaReq.Name,
			Description: areaReq.Description,
		}
		building.Areas = append(building.Areas, area)
	}

	return s.buildingRepo.CreateBuilding(&building)
}

func (s *BuildingService) GetAllBuildings(accountID uint) ([]entities.Building, error) {
	return s.buildingRepo.FindAllByAccountID(accountID)
}
