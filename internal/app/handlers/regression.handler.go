package handlers

import (
	"core-api/internal/app/dto"
	"core-api/internal/domain/services"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// RegressionHandler handles HTTP requests for regression operations.
type RegressionHandler struct {
	service services.RegressionService
}

// NewRegressionHandler creates a new instance of RegressionHandler.
func NewRegressionHandler(service services.RegressionService) *RegressionHandler {
	return &RegressionHandler{
		service: service,
	}
}

// CreateRegression handles POST /regressions.
func (h *RegressionHandler) CreateRegression(c *gin.Context) {
	var req dto.CreateRegressionRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	regression, err := h.service.CreateRegression(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create regression"})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"message": "Regression created successfully", "regression": regression})
}

// ListRegressions handles GET /regressions.
func (h *RegressionHandler) ListRegressions(c *gin.Context) {
	regressions, err := h.service.ListRegressions()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to list regressions"})
		return
	}
	c.JSON(http.StatusOK, regressions)
}

// GetRegressionByID handles GET /regressions/:id.
func (h *RegressionHandler) GetRegressionByID(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid regression ID"})
		return
	}
	regression, err := h.service.GetRegressionByID(uint(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve regression"})
		return
	}
	if regression == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Regression not found"})
		return
	}
	c.JSON(http.StatusOK, regression)
}

// UpdateRegression handles PUT /regressions/:id.
func (h *RegressionHandler) UpdateRegression(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid regression ID"})
		return
	}
	var req dto.UpdateRegressionRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := h.service.UpdateRegressionByID(uint(id), req); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update regression"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Regression updated successfully"})
}

// DeleteRegression handles DELETE /regressions/:id.
func (h *RegressionHandler) DeleteRegression(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid regression ID"})
		return
	}
	if err := h.service.DeleteRegressionByID(uint(id)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete regression"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Regression deleted successfully"})
}
