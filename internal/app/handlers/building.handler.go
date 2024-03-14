package handlers

import (
	"core-api/internal/app/dto"
	"core-api/internal/domain/entities"
	"core-api/internal/domain/services"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"reflect"
	"strconv"
)

type BuildingHandler struct {
	accountService  *services.AccountService
	buildingService *services.BuildingService
}

func NewBuildingHandler(accountService *services.AccountService, buildingService *services.BuildingService) *BuildingHandler {
	return &BuildingHandler{
		accountService:  accountService,
		buildingService: buildingService,
	}
}

// HandleCreateBuilding processes the POST request to create a new building with areas.
func (h *BuildingHandler) HandleCreateBuilding(c *gin.Context) {
	var req dto.CreateBuildingRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request data"})
		return
	}

	accountIDInterface, exists := c.Get("accountId")
	if !exists {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Account ID is required"})
		return
	}

	accountIDFloat, ok := accountIDInterface.(float64)
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Account ID must be a number"})
		return
	}

	// Call the service to handle the business logic for building creation.
	if err := h.buildingService.CreateBuilding(&req, uint(accountIDFloat)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create building"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Building created successfully", "building": req})
}

func (h *BuildingHandler) GetAllBuildings(c *gin.Context) {
	accountIDInterface, exists := c.Get("accountId")
	if !exists {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Account ID is required"})
		return
	}

	fmt.Println(reflect.TypeOf(accountIDInterface))
	accountID, ok := accountIDInterface.(float64)
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Account ID is invalid"})
		return
	}

	buildings, err := h.buildingService.GetAllBuildings(uint(accountID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve buildings"})
		return
	}

	c.JSON(http.StatusOK, buildings)
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
