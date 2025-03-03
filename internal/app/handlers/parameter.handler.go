package handlers

import (
	"core-api/internal/app/dto"
	"core-api/internal/domain/services"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type ParameterHandler struct {
	service services.ParameterService
}

func NewParameterHandler(service services.ParameterService) *ParameterHandler {
	return &ParameterHandler{
		service: service,
	}
}

// CreateParameter handles POST /parameter.
func (h *ParameterHandler) CreateParameter(c *gin.Context) {
	var req dto.CreateParameterRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	param, err := h.service.CreateParameter(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create parameter: " + err.Error()})
		return
	}
	c.JSON(http.StatusCreated, param)
}

// GetParameterByID handles GET /parameter/:id.
func (h *ParameterHandler) GetParameterByID(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid parameter ID"})
		return
	}
	param, err := h.service.GetParameterByID(uint(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve parameter: " + err.Error()})
		return
	}
	c.JSON(http.StatusOK, param)
}

// UpdateParameter handles PUT /parameter/:id.
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
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update parameter: " + err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Parameter updated successfully"})
}
