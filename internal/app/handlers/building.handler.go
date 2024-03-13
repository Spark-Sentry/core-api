package handlers

import (
	"core-api/internal/app/dto"
	"core-api/internal/domain/services"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"reflect"
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
