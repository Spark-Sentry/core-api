package handlers

import (
	"core-api/internal/app/dto"
	"core-api/internal/domain/services"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// SubsidyHandler handles HTTP requests for subsidy operations.
type SubsidyHandler struct {
	service services.SubsidyService
}

// NewSubsidyHandler creates a new instance of SubsidyHandler.
func NewSubsidyHandler(service services.SubsidyService) *SubsidyHandler {
	return &SubsidyHandler{
		service: service,
	}
}

// CreateSubsidy handles POST /subsidies.
func (h *SubsidyHandler) CreateSubsidy(c *gin.Context) {
	var req dto.CreateSubsidyRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	subsidy, err := h.service.CreateSubsidy(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create subsidy: " + err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"message": "Subsidy created successfully", "subsidy": subsidy})
}

// ListSubsidies handles GET /subsidies.
func (h *SubsidyHandler) ListSubsidies(c *gin.Context) {
	subsidies, err := h.service.ListSubsidies()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to list subsidies"})
		return
	}
	c.JSON(http.StatusOK, subsidies)
}

// GetSubsidyByID handles GET /subsidies/:id.
func (h *SubsidyHandler) GetSubsidyByID(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid subsidy ID"})
		return
	}
	subsidy, err := h.service.GetSubsidyByID(uint(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve subsidy"})
		return
	}
	if subsidy == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Subsidy not found"})
		return
	}
	c.JSON(http.StatusOK, subsidy)
}

// UpdateSubsidy handles PUT /subsidies/:id.
func (h *SubsidyHandler) UpdateSubsidy(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid subsidy ID"})
		return
	}
	var req dto.UpdateSubsidyRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := h.service.UpdateSubsidyByID(uint(id), req); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update subsidy"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Subsidy updated successfully"})
}

// DeleteSubsidy handles DELETE /subsidies/:id.
func (h *SubsidyHandler) DeleteSubsidy(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid subsidy ID"})
		return
	}
	if err := h.service.DeleteSubsidyByID(uint(id)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete subsidy"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Subsidy deleted successfully"})
}
