package handlers

import (
	"core-api/internal/app/dto"
	"core-api/internal/domain/entities"
	"core-api/internal/domain/services"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type BuildingHandler struct {
	accountService  *services.AccountService
	buildingService services.BuildingService
}

func NewBuildingHandler(accountService *services.AccountService, buildingService *services.BuildingService) *BuildingHandler {
	return &BuildingHandler{
		accountService:  accountService,
		buildingService: *buildingService,
	}
}

// HandleCreateBuilding processes the POST request to create a new building with areas.
func (h *BuildingHandler) HandleCreateBuilding(c *gin.Context) {
	var req dto.CreateBuildingRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request data"})
		return
	}

	user, exists := c.Get("user")
	if !exists {
		c.JSON(http.StatusBadRequest, gin.H{"error": "User not found in context"})
		return
	}
	userDetails := user.(*entities.User)

	if err := h.buildingService.CreateBuilding(&req, *userDetails.AccountID); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create building"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Building created successfully", "building": req})
}

func (h *BuildingHandler) GetAllBuildings(c *gin.Context) {
	user, exists := c.Get("user")
	if !exists {
		c.JSON(http.StatusBadRequest, gin.H{"error": "User not found in context"})
		return
	}
	userDetails := user.(*entities.User)

	buildings, err := h.buildingService.GetAllBuildings(*userDetails.AccountID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve buildings"})
		return
	}

	c.JSON(http.StatusOK, buildings)
}

// AddArea handles the request to add one or more new areas to their respective buildings.
func (h *BuildingHandler) AddArea(c *gin.Context) {
	var requestDTO dto.AreaCreateDTO
	buildingID, err := strconv.ParseUint(c.Param("building_id"), 10, 32)
	if err := c.ShouldBindJSON(&requestDTO); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	createdAreas, err := h.buildingService.AddAreas(requestDTO.Areas, uint(buildingID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"data": createdAreas})
}

// GetAreasByBuildingID handles the request to retrieve areas for a specific building
func (h *BuildingHandler) GetAreasByBuildingID(c *gin.Context) {
	buildingIdParam := c.Param("building_id")
	buildingID, err := strconv.ParseUint(buildingIdParam, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid building ID"})
		return
	}

	areas, err := h.buildingService.GetAreasByBuildingID(uint(buildingID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, areas)
}

// UpdateArea handles the PUT request to update an area's details.
func (h *BuildingHandler) UpdateArea(c *gin.Context) {
	areaIDParam := c.Param("area_id")
	areaID, err := strconv.ParseUint(areaIDParam, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid area ID"})
		return
	}

	var updateDTO dto.AreaUpdateDTO
	if err := c.ShouldBindJSON(&updateDTO); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.buildingService.UpdateArea(uint(areaID), updateDTO); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update area"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Area updated successfully"})
}

// DeleteArea handles the DELETE request to delete a specific area.
func (h *BuildingHandler) DeleteArea(c *gin.Context) {
	areaIDParam := c.Param("area_id") // Extract the area ID from the URL
	areaID, err := strconv.ParseUint(areaIDParam, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid area ID"})
		return
	}

	if err := h.buildingService.DeleteArea(uint(areaID)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete area"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Area deleted successfully"})
}

func (h *BuildingHandler) AddSystem(c *gin.Context) {
	areaID, _ := strconv.Atoi(c.Param("area_id"))
	var systemData entities.System
	if err := c.ShouldBindJSON(&systemData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := h.buildingService.AddSystemToArea(uint(areaID), &systemData)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to add system to building"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "System added successfully"})
}

func (h *BuildingHandler) GetSystemsByAreaID(c *gin.Context) {
	areaID, err := strconv.Atoi(c.Param("area_id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid area ID"})
		return
	}

	systems, err := h.buildingService.GetSystemsByAreaID(uint(areaID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve systems"})
		return
	}

	c.JSON(http.StatusOK, systems)
}

func (h *BuildingHandler) UpdateSystem(c *gin.Context) {
	systemID, err := strconv.Atoi(c.Param("system_id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid system ID"})
		return
	}

	var updateDTO dto.SystemUpdateDTO
	if err := c.ShouldBindJSON(&updateDTO); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.buildingService.UpdateSystem(uint(systemID), updateDTO); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update system"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "System updated successfully"})
}

func (h *BuildingHandler) DeleteSystem(c *gin.Context) {
	systemID, err := strconv.Atoi(c.Param("system_id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid system ID"})
		return
	}

	if err := h.buildingService.DeleteSystem(uint(systemID)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete system"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "System removed successfully"})
}

// AddEquipmentToSystem handles the POST request to add new equipment to a system.
func (h *BuildingHandler) AddEquipmentToSystem(c *gin.Context) {
	systemID, err := strconv.ParseUint(c.Param("system_id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid system ID"})
		return
	}

	var equipmentDTOs []dto.EquipmentCreateDTO
	if err := c.ShouldBindJSON(&equipmentDTOs); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	for _, equipmentDTO := range equipmentDTOs {
		err := h.buildingService.AddEquipmentToSystem(uint(systemID), equipmentDTO)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to add equipment to system"})
			return
		}
	}

	c.JSON(http.StatusOK, gin.H{"message": "Equipment added successfully"})
}

// GetEquipmentsBySystemID handles the GET request to list all equipments within a specific system.
func (h *BuildingHandler) GetEquipmentsBySystemID(c *gin.Context) {
	systemID, err := strconv.Atoi(c.Param("system_id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid system ID"})
		return
	}

	equipments, err := h.buildingService.GetEquipmentsBySystemID(uint(systemID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve equipments"})
		return
	}

	c.JSON(http.StatusOK, equipments)
}

// UpdateEquipment handles the PUT request to update a specific piece of equipment.
func (h *BuildingHandler) UpdateEquipment(c *gin.Context) {
	equipmentID, err := strconv.Atoi(c.Param("equipment_id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid equipmentID"})
		return
	}

	var updateDTO dto.EquipmentUpdateDTO
	if err := c.ShouldBindJSON(&updateDTO); err != nil { // Bind the JSON body to DTO
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.buildingService.UpdateEquipment(uint(equipmentID), updateDTO); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update equipment"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Equipment updated successfully"})
}

// DeleteEquipment handles the DELETE request to remove a specific piece of equipment.
func (h *BuildingHandler) DeleteEquipment(c *gin.Context) {
	equipmentIDParam := c.Param("equipment_id") // Extract the equipment ID from URL
	equipmentID, err := strconv.ParseUint(equipmentIDParam, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid equipment ID"})
		return
	}

	if err := h.buildingService.DeleteEquipment(uint(equipmentID)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to remove equipment"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Equipment removed successfully"})
}

// AddParameterToEquipment handles the POST request to add new parameters to a specific piece of equipment.
func (h *BuildingHandler) AddParameterToEquipment(c *gin.Context) {
	equipmentID, err := strconv.ParseUint(c.Param("equipment_id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid equipment ID"})
		return
	}

	var parameterDTOs []dto.ParameterCreateDTO
	if err := c.ShouldBindJSON(&parameterDTOs); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	for _, parameterDTO := range parameterDTOs {
		err := h.buildingService.AddParameterToEquipment(uint(equipmentID), parameterDTO)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to add parameter to equipment"})
			return
		}
	}

	c.JSON(http.StatusOK, gin.H{"message": "Parameters added successfully"})
}

// GetParametersByEquipmentID handles the GET request to list all parameters within a specific piece of equipment.
func (h *BuildingHandler) GetParametersByEquipmentID(c *gin.Context) {
	equipmentID, err := strconv.Atoi(c.Param("equipment_id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid equipment ID"})
		return
	}

	parameters, err := h.buildingService.GetParametersByEquipmentID(uint(equipmentID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve parameters"})
	}
	c.JSON(http.StatusOK, parameters)
}
