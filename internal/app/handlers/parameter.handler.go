package handlers

import (
	"core-api/internal/app/dto"
	"core-api/internal/domain/services"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// ParameterHandler handles HTTP requests for parameter management.
type ParameterHandler struct {
	service services.ParameterService
}

// NewParameterHandler creates a new instance of ParameterHandler.
func NewParameterHandler(service services.ParameterService) *ParameterHandler {
	return &ParameterHandler{
		service: service,
	}
}

// CreateParameter handles POST /parameters.
func (h *ParameterHandler) CreateParameter(c *gin.Context) {
	var req dto.CreateParameterRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	param, err := h.service.CreateParameter(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create parameter"})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"message": "Parameter created successfully", "parameter": param})
}

// ListParameters handles GET /parameters.
func (h *ParameterHandler) ListParameters(c *gin.Context) {
	params, err := h.service.ListParameters()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to list parameters"})
		return
	}
	c.JSON(http.StatusOK, params)
}

// GetParameterByID handles GET /parameters/:id.
func (h *ParameterHandler) GetParameterByID(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid parameter ID"})
		return
	}
	param, err := h.service.GetParameterByID(uint(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve parameter"})
		return
	}
	if param == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Parameter not found"})
		return
	}
	c.JSON(http.StatusOK, param)
}

// UpdateParameter handles PUT /parameters/:id.
func (h *ParameterHandler) UpdateParameter(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid parameter ID"})
		return
	}
	var req dto.UpdateParameterRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := h.service.UpdateParameterByID(uint(id), req); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update parameter"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Parameter updated successfully"})
}

// DeleteParameter handles DELETE /parameters/:id.
func (h *ParameterHandler) DeleteParameter(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid parameter ID"})
		return
	}
	if err := h.service.DeleteParameterByID(uint(id)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete parameter"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Parameter deleted successfully"})
}
