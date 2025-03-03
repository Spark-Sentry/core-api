package handlers

import (
	"core-api/internal/app/dto"
	"core-api/internal/domain/services"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// TargetHandler handles HTTP requests for target management.
type TargetHandler struct {
	service services.TargetService
}

// NewTargetHandler creates a new instance of TargetHandler.
func NewTargetHandler(service services.TargetService) *TargetHandler {
	return &TargetHandler{
		service: service,
	}
}

// CreateTarget handles POST /targets.
func (h *TargetHandler) CreateTarget(c *gin.Context) {
	var req dto.CreateTargetRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	target, err := h.service.CreateTarget(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create target"})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"message": "Target created successfully", "target": target})
}

// ListTargets handles GET /targets.
func (h *TargetHandler) ListTargets(c *gin.Context) {
	targets, err := h.service.ListTargets()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to list targets"})
		return
	}
	c.JSON(http.StatusOK, targets)
}

// GetTargetByID handles GET /targets/:id.
func (h *TargetHandler) GetTargetByID(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid target ID"})
		return
	}
	target, err := h.service.GetTargetByID(uint(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve target"})
		return
	}
	if target == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Target not found"})
		return
	}
	c.JSON(http.StatusOK, target)
}

// UpdateTarget handles PUT /targets/:id.
func (h *TargetHandler) UpdateTarget(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid target ID"})
		return
	}
	var req dto.UpdateTargetRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := h.service.UpdateTargetByID(uint(id), req); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update target"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Target updated successfully"})
}

// DeleteTarget handles DELETE /targets/:id.
func (h *TargetHandler) DeleteTarget(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid target ID"})
		return
	}
	if err := h.service.DeleteTargetByID(uint(id)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete target"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Target deleted successfully"})
}
