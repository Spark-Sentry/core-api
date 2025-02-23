package handlers

import (
	"core-api/internal/app/dto"
	"core-api/internal/domain/entities"
	"core-api/internal/domain/services"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// EfficiencyMeasureHandler handles HTTP requests related to efficiency measures.
type EfficiencyMeasureHandler struct {
	service services.EfficiencyMeasureService
}

// NewEfficiencyMeasureHandler creates a new instance of EfficiencyMeasureHandler.
func NewEfficiencyMeasureHandler(service services.EfficiencyMeasureService) *EfficiencyMeasureHandler {
	return &EfficiencyMeasureHandler{
		service: service,
	}
}

// CreateEfficiencyMeasure handles POST /efficiencymeasures.
func (h *EfficiencyMeasureHandler) CreateEfficiencyMeasure(c *gin.Context) {
	var req dto.CreateEfficiencyMeasureRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	measure := entities.EfficiencyMeasure{
		ProjectID:       req.ProjectID,
		Name:            req.Name,
		EfficiencyCost:  req.EfficiencyCost,
		MaintenanceCost: req.MaintenanceCost,
	}
	if err := h.service.CreateEfficiencyMeasure(&measure); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create efficiency measure"})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"message": "Efficiency measure created successfully", "efficiencyMeasure": measure})
}

// ListEfficiencyMeasures handles GET /efficiencymeasures.
func (h *EfficiencyMeasureHandler) ListEfficiencyMeasures(c *gin.Context) {
	measures, err := h.service.ListEfficiencyMeasures()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to list efficiency measures"})
		return
	}
	c.JSON(http.StatusOK, measures)
}

// GetEfficiencyMeasureByID handles GET /efficiencymeasures/:id.
func (h *EfficiencyMeasureHandler) GetEfficiencyMeasureByID(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid efficiency measure ID"})
		return
	}
	measure, err := h.service.GetEfficiencyMeasureByID(uint(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve efficiency measure"})
		return
	}
	if measure == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Efficiency measure not found"})
		return
	}
	c.JSON(http.StatusOK, measure)
}

// UpdateEfficiencyMeasure handles PUT /efficiencymeasures/:id.
func (h *EfficiencyMeasureHandler) UpdateEfficiencyMeasure(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid efficiency measure ID"})
		return
	}
	var req dto.UpdateEfficiencyMeasureRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := h.service.UpdateEfficiencyMeasureByID(uint(id), req); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update efficiency measure"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Efficiency measure updated successfully"})
}
