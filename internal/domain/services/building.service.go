package services

import (
	"core-api/internal/app/dto"
	"core-api/internal/domain/entities"
	"core-api/internal/infrastructure/repository"
	"errors"
	"fmt"
)

// BuildingService interface for building-related services.
type BuildingService interface {
	CreateBuilding(req *dto.CreateBuildingRequest, accountID uint) error
	GetAllBuildings(accountID uint) ([]entities.Building, error)
	AddAreas(areasDTO []dto.AreaDTO, buildingID uint) ([]entities.Area, error)
	AddSystemToArea(areaID uint, systemData *entities.System) error
	GetSystemsByAreaID(areaID uint) ([]entities.System, error)
	AddEquipmentToSystem(systemID uint, equipmentDTO dto.EquipmentCreateDTO) error
	GetAreasByBuildingID(buildingID uint) ([]entities.Area, error)
	UpdateArea(areaID uint, dto dto.AreaUpdateDTO) error
	DeleteArea(areaID uint) error
	UpdateSystem(systemID uint, dto dto.SystemUpdateDTO) error
	DeleteSystem(systemID uint) error
	GetEquipmentsBySystemID(systemID uint) ([]entities.Equipment, error)
	UpdateEquipment(equipmentID uint, dto dto.EquipmentUpdateDTO) error
	DeleteEquipment(equipmentID uint) error
	AddParameterToEquipment(equipmentID uint, parameterDTO dto.ParameterCreateDTO) error
	GetParametersByEquipmentID(equipmentID uint) ([]entities.Parameter, error)
	UpdateParameter(parameterID uint, dto dto.ParameterUpdateDTO) error
	DeleteParameter(parameterID uint) error
}

type buildingService struct {
	areaRepo      repository.AreaRepository
	buildingRepo  repository.BuildingRepository
	systemRepo    repository.SystemRepository
	equipmentRepo repository.EquipmentRepository
	parameterRepo repository.ParameterRepository
}

// NewBuildingService creates a new instance of BuildingService.
func NewBuildingService(
	buildingRepo repository.BuildingRepository,
	systemRepo repository.SystemRepository,
	equipmentRepo repository.EquipmentRepository,
	areaRepo repository.AreaRepository,
	parameterRepo repository.ParameterRepository,
) BuildingService {
	return &buildingService{
		buildingRepo:  buildingRepo,
		systemRepo:    systemRepo,
		equipmentRepo: equipmentRepo,
		areaRepo:      areaRepo,
		parameterRepo: parameterRepo,
	}
}

// CreateBuilding creates a new Building with at least one Area from a DTO.
func (s *buildingService) CreateBuilding(req *dto.CreateBuildingRequest, accountID uint) error {
	if req == nil {
		return errors.New("invalid request: request body is nil")
	}

	if len(req.Areas) == 0 {
		return errors.New("a building must have at least one area")
	}

	building := entities.Building{
		Name:      req.Name,
		Address:   req.Address,
		Group:     req.Group,
		AccountID: accountID,
		Areas:     []entities.Area{},
	}

	for _, areaReq := range req.Areas {
		building.Areas = append(building.Areas, entities.Area{
			Name:        areaReq.Name,
			Description: areaReq.Description,
		})
	}

	err := s.buildingRepo.CreateBuilding(&building)
	if err != nil {
		return fmt.Errorf("failed to create building: %w", err)
	}
	return nil
}

func (s *buildingService) GetAllBuildings(accountID uint) ([]entities.Building, error) {
	buildings, err := s.buildingRepo.FindAllByAccountID(accountID)
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve buildings: %w", err)
	}
	return buildings, nil
}

func (s *buildingService) AddAreas(areasDTO []dto.AreaDTO, buildingID uint) ([]entities.Area, error) {
	if len(areasDTO) == 0 {
		return nil, errors.New("no areas provided to add")
	}

	var areas []entities.Area
	for _, areaDTO := range areasDTO {
		area := entities.Area{
			BuildingID:  buildingID,
			Name:        areaDTO.Name,
			Description: areaDTO.Description,
		}
		createdArea, err := s.areaRepo.Save(&area)
		if err != nil {
			return nil, fmt.Errorf("failed to add area: %w", err)
		}
		areas = append(areas, *createdArea)
	}
	return areas, nil
}

func (s *buildingService) GetAreasByBuildingID(buildingID uint) ([]entities.Area, error) {
	areas, err := s.areaRepo.FindByBuildingID(buildingID)
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve areas: %w", err)
	}
	return areas, nil
}

func (s *buildingService) UpdateArea(areaID uint, dto dto.AreaUpdateDTO) error {
	if dto.Name == "" {
		return errors.New("area name cannot be empty")
	}

	area := entities.Area{
		ID:          areaID,
		Name:        dto.Name,
		Description: dto.Description,
	}

	err := s.areaRepo.UpdateArea(area)
	if err != nil {
		return fmt.Errorf("failed to update area: %w", err)
	}
	return nil
}

func (s *buildingService) DeleteArea(areaID uint) error {
	err := s.areaRepo.DeleteArea(areaID)
	if err != nil {
		return fmt.Errorf("failed to delete area: %w", err)
	}
	return nil
}

func (s *buildingService) AddSystemToArea(areaID uint, systemData *entities.System) error {
	if systemData.Name == "" {
		return errors.New("system name cannot be empty")
	}

	system := &entities.System{
		AreaID:      areaID,
		Name:        systemData.Name,
		Description: systemData.Description,
	}

	err := s.systemRepo.CreateSystem(system)
	if err != nil {
		return fmt.Errorf("failed to add system to area: %w", err)
	}
	return nil
}

func (s *buildingService) GetSystemsByAreaID(areaID uint) ([]entities.System, error) {
	systems, err := s.systemRepo.FindByAreaID(areaID)
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve systems for area ID %d: %w", areaID, err)
	}
	return systems, nil
}

func (s *buildingService) UpdateSystem(systemID uint, dto dto.SystemUpdateDTO) error {
	return s.systemRepo.UpdateSystem(systemID, dto)
}

func (s *buildingService) DeleteSystem(systemID uint) error {
	return s.systemRepo.DeleteSystem(systemID)
}

func (s *buildingService) AddEquipmentToSystem(systemID uint, equipmentDTO dto.EquipmentCreateDTO) error {
	if equipmentDTO.Tag == "" {
		return errors.New("equipment tag cannot be empty")
	}

	equipment := entities.Equipment{
		SystemID:    systemID,
		Tag:         equipmentDTO.Tag,
		Description: equipmentDTO.Description,
	}

	err := s.equipmentRepo.AddEquipment(&equipment)
	if err != nil {
		return fmt.Errorf("failed to add equipment to system: %w", err)
	}
	return nil
}

func (s *buildingService) GetEquipmentsBySystemID(systemID uint) ([]entities.Equipment, error) {
	equipments, err := s.equipmentRepo.FindBySystemID(systemID)
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve equipment for system ID %d: %w", systemID, err)
	}
	return equipments, nil
}

func (s *buildingService) UpdateEquipment(equipmentID uint, dto dto.EquipmentUpdateDTO) error {
	err := s.equipmentRepo.UpdateEquipment(equipmentID, dto)
	if err != nil {
		return fmt.Errorf("failed to update equipment: %w", err)
	}
	return nil
}

func (s *buildingService) DeleteEquipment(equipmentID uint) error {
	err := s.equipmentRepo.DeleteEquipment(equipmentID)
	if err != nil {
		return fmt.Errorf("failed to delete equipment: %w", err)
	}
	return nil
}

func (s *buildingService) AddParameterToEquipment(equipmentID uint, parameterDTO dto.ParameterCreateDTO) error {
	if parameterDTO.Name == "" {
		return errors.New("parameter name cannot be empty")
	}

	parameter := entities.Parameter{
		Name:        parameterDTO.Name,
		HostDevice:  parameterDTO.HostDevice,
		Device:      parameterDTO.Device,
		Log:         parameterDTO.Log,
		Point:       parameterDTO.Point,
		Unit:        parameterDTO.Unit,
		EquipmentID: equipmentID,
	}

	err := s.parameterRepo.Create(&parameter)
	if err != nil {
		return fmt.Errorf("failed to add parameter to equipment: %w", err)
	}
	return nil
}

func (s *buildingService) GetParametersByEquipmentID(equipmentID uint) ([]entities.Parameter, error) {
	parameters, err := s.parameterRepo.FindByEquipmentID(equipmentID)
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve parameters for equipment ID %d: %w", equipmentID, err)
	}
	return parameters, nil
}

func (s *buildingService) UpdateParameter(parameterID uint, dto dto.ParameterUpdateDTO) error {
	err := s.parameterRepo.Update(&entities.Parameter{
		ID:         parameterID,
		Name:       dto.Name,
		HostDevice: dto.HostDevice,
		Device:     dto.Device,
		Log:        dto.Log,
		Point:      dto.Point,
		Unit:       dto.Unit,
	})
	if err != nil {
		return fmt.Errorf("failed to update parameter: %w", err)
	}
	return nil
}

func (s *buildingService) DeleteParameter(parameterID uint) error {
	err := s.parameterRepo.Delete(parameterID)
	if err != nil {
		return fmt.Errorf("failed to delete parameter: %w", err)
	}
	return nil
}
