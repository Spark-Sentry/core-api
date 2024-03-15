package services

import (
	"core-api/internal/app/dto"
	"core-api/internal/domain/entities"
	"core-api/internal/infrastructure/repository"
)

// BuildingService interface for building related services.
type BuildingService interface {
	// CreateBuilding creates a new Building with at least one Area.
	CreateBuilding(req *dto.CreateBuildingRequest, accountID uint) error

	// GetAllBuildings retrieves all buildings associated with an account.
	GetAllBuildings(accountID uint) ([]entities.Building, error)

	// AddAreas adds a new area to a building.
	AddAreas(areasDTO []dto.AreaDTO, buildingID uint) ([]entities.Area, error)

	// AddSystemToArea adds a new system to an area.
	AddSystemToArea(areaID uint, systemData *entities.System) error

	// GetSystemsByAreaID retrieves all systems associated with an area.
	GetSystemsByAreaID(areaID uint) ([]entities.System, error)

	// AddEquipmentToSystem adds new equipment to a system.
	AddEquipmentToSystem(systemID uint, equipmentDTO dto.EquipmentCreateDTO) error

	// GetAreasByBuildingID retrieves areas for a specific building
	GetAreasByBuildingID(buildingID uint) ([]entities.Area, error)

	// UpdateArea updates the specified area with new details.
	UpdateArea(areaID uint, dto dto.AreaUpdateDTO) error

	// DeleteArea deletes the specified area.
	DeleteArea(areaID uint) error

	// UpdateSystem update the specified system
	UpdateSystem(systemID uint, dto dto.SystemUpdateDTO) error

	// DeleteSystem delete the specified system
	DeleteSystem(systemID uint) error

	// GetEquipmentsBySystemID retrieves all equipments within a specific system.
	GetEquipmentsBySystemID(systemID uint) ([]entities.Equipment, error)

	// UpdateEquipment updates the specified piece of equipment with new details.
	UpdateEquipment(equipmentID uint, dto dto.EquipmentUpdateDTO) error

	// DeleteEquipment deletes the specified piece of equipment.
	DeleteEquipment(equipmentID uint) error
}

type buildingService struct {
	areaRepo      repository.AreaRepository
	buildingRepo  repository.BuildingRepository
	systemRepo    repository.SystemRepository
	equipmentRepo repository.EquipmentRepository
}

// NewBuildingService creates a new instance of BuildingService.
func NewBuildingService(buildingRepo repository.BuildingRepository, systemRepo repository.SystemRepository, equipmentRepo repository.EquipmentRepository, areaRepo repository.AreaRepository) BuildingService {
	return &buildingService{buildingRepo: buildingRepo, systemRepo: systemRepo, equipmentRepo: equipmentRepo, areaRepo: areaRepo}
}

// CreateBuilding creates a new Building with at least one Area from a DTO.
func (s *buildingService) CreateBuilding(req *dto.CreateBuildingRequest, accountID uint) error {
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

func (s *buildingService) GetAllBuildings(accountID uint) ([]entities.Building, error) {
	return s.buildingRepo.FindAllByAccountID(accountID)
}

// AddAreas iterates over the slice of AreaDTO, adds each Area, and returns a slice of the added Areas
func (s *buildingService) AddAreas(areasDTO []dto.AreaDTO, buildingID uint) ([]entities.Area, error) {
	var areas []entities.Area
	for _, areaDTO := range areasDTO {
		area := entities.Area{
			BuildingID:  buildingID,
			Name:        areaDTO.Name,
			Description: areaDTO.Description,
		}
		createdArea, err := s.areaRepo.Save(&area)
		if err != nil {
			return nil, err
		}
		areas = append(areas, *createdArea)
	}
	return areas, nil
}

// GetAreasByBuildingID retrieves areas for a specific building
func (s *buildingService) GetAreasByBuildingID(buildingID uint) ([]entities.Area, error) {
	return s.areaRepo.FindByBuildingID(buildingID)
}

// UpdateArea updates the specified area with new details.
func (s *buildingService) UpdateArea(areaID uint, dto dto.AreaUpdateDTO) error {
	area := entities.Area{ID: areaID, Name: dto.Name, Description: dto.Description}
	return s.areaRepo.UpdateArea(area)
}

// DeleteArea deletes the specified area.
func (s *buildingService) DeleteArea(areaID uint) error {
	return s.areaRepo.DeleteArea(areaID)
}

func (s *buildingService) AddSystemToArea(areaID uint, systemData *entities.System) error {
	system := &entities.System{
		AreaID:      areaID,
		Name:        systemData.Name,
		Description: systemData.Description,
	}
	return s.systemRepo.CreateSystem(system)
}

func (s *buildingService) GetSystemsByAreaID(areaID uint) ([]entities.System, error) {
	return s.systemRepo.FindByAreaID(areaID)
}

func (s *buildingService) UpdateSystem(systemID uint, dto dto.SystemUpdateDTO) error {
	return s.systemRepo.UpdateSystem(systemID, dto)
}

func (s *buildingService) DeleteSystem(systemID uint) error {
	return s.systemRepo.DeleteSystem(systemID)
}

func (s *buildingService) AddEquipmentToSystem(systemID uint, equipmentDTO dto.EquipmentCreateDTO) error {
	equipment := entities.Equipment{
		SystemID:    systemID,
		Tag:         equipmentDTO.Tag,
		Description: equipmentDTO.Description,
	}
	return s.equipmentRepo.AddEquipment(&equipment)
}

// GetEquipmentsBySystemID retrieves all equipments within a specific system.
func (s *buildingService) GetEquipmentsBySystemID(systemID uint) ([]entities.Equipment, error) {
	return s.equipmentRepo.FindBySystemID(systemID)
}

// UpdateEquipment updates the specified piece of equipment with new details.
func (s *buildingService) UpdateEquipment(equipmentID uint, dto dto.EquipmentUpdateDTO) error {
	return s.equipmentRepo.UpdateEquipment(equipmentID, dto)
}

// DeleteEquipment deletes the specified piece of equipment.
func (s *buildingService) DeleteEquipment(equipmentID uint) error {
	return s.equipmentRepo.DeleteEquipment(equipmentID)
}
