package handlers

import (
	"core-api/internal/app/dto"
	"core-api/internal/domain/services"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// IndependantVariableHandler handles HTTP requests for independant variable management.
type IndependantVariableHandler struct {
	service services.IndependantVariableService
}

// NewIndependantVariableHandler creates a new instance of IndependantVariableHandler.
func NewIndependantVariableHandler(service services.IndependantVariableService) *IndependantVariableHandler {
	return &IndependantVariableHandler{
		service: service,
	}
}

// CreateIndependantVariable handles POST /independantvariables.
func (h *IndependantVariableHandler) CreateIndependantVariable(c *gin.Context) {
	var req dto.CreateIndependantVariableRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	iv, err := h.service.CreateIndependantVariable(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create independant variable"})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"message": "Independant variable created successfully", "independantVariable": iv})
}

// ListIndependantVariables handles GET /independantvariables.
func (h *IndependantVariableHandler) ListIndependantVariables(c *gin.Context) {
	ivs, err := h.service.ListIndependantVariables()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to list independant variables"})
		return
	}
	c.JSON(http.StatusOK, ivs)
}

// GetIndependantVariableByID handles GET /independantvariables/:id.
func (h *IndependantVariableHandler) GetIndependantVariableByID(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid independant variable ID"})
		return
	}
	iv, err := h.service.GetIndependantVariableByID(uint(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve independant variable"})
		return
	}
	if iv == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Independant variable not found"})
		return
	}
	c.JSON(http.StatusOK, iv)
}

// UpdateIndependantVariable handles PUT /independantvariables/:id.
func (h *IndependantVariableHandler) UpdateIndependantVariable(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid independant variable ID"})
		return
	}
	var req dto.UpdateIndependantVariableRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := h.service.UpdateIndependantVariableByID(uint(id), req); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update independant variable"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Independant variable updated successfully"})
}

// DeleteIndependantVariable handles DELETE /independantvariables/:id.
func (h *IndependantVariableHandler) DeleteIndependantVariable(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid independant variable ID"})
		return
	}
	if err := h.service.DeleteIndependantVariableByID(uint(id)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete independant variable"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Independant variable deleted successfully"})
}
