package handlers

import (
	"core-api/internal/app/dto"
	"core-api/internal/domain/entities"
	"core-api/internal/domain/services"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// BuildingHandler handles building-related HTTP requests.
type BuildingHandler struct {
	buildingService services.BuildingService
}

// NewBuildingHandler creates a new instance of BuildingHandler.
func NewBuildingHandler(buildingService services.BuildingService) *BuildingHandler {
	return &BuildingHandler{
		buildingService: buildingService,
	}
}

// CreateBuilding handles the POST request to create a new building.
// CHANGES:
// - Uses new DTO fields: Name, Address, and CategoryID.
// - Converts req.CategoryID (uint) to a pointer (*uint).
func (h *BuildingHandler) CreateBuilding(c *gin.Context) {
	var req dto.CreateBuildingRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request data"})
		return
	}

	// Retrieve the authenticated user from context.
	user, exists := c.Get("user")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not authenticated"})
		return
	}
	userDetails := user.(*entities.User)

	// Convert req.CategoryID (uint) to a pointer.
	categoryIDPtr := &req.CategoryID

	// Create a Building instance.
	building := entities.Building{
		AccountID:  *userDetails.AccountID,
		Name:       req.Name,
		Address:    req.Address,
		CategoryID: categoryIDPtr,
	}

	if err := h.buildingService.CreateBuilding(&building); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create building"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Building created successfully", "building": building})
}

// GetAllBuildings handles the GET request to retrieve all buildings for an account.
func (h *BuildingHandler) GetAllBuildings(c *gin.Context) {
	// Retrieve the authenticated user from context.
	user, exists := c.Get("user")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not authenticated"})
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

// UpdateBuilding handles PUT /buildings/:id.
func (h *BuildingHandler) UpdateBuilding(c *gin.Context) {
	// Get the building ID from the URL.
	idParam := c.Param("id")
	buildingID, err := strconv.ParseUint(idParam, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid building ID"})
		return
	}

	var req dto.UpdateBuildingRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Retrieve the authenticated user (to get accountId)
	user, exists := c.Get("user")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not authenticated"})
		return
	}
	userDetails := user.(*entities.User)

	// Build the updated building data.
	updatedBuilding := entities.Building{
		Name:       req.Name,
		Address:    req.Address,
		CategoryID: &req.CategoryID,
	}

	// Call the service to update the building.
	if err := h.buildingService.UpdateBuilding(uint(buildingID), updatedBuilding, *userDetails.AccountID); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update building: " + err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Building updated successfully"})
}
