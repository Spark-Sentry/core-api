package handlers

import (
	"core-api/internal/app/dto"
	"core-api/internal/domain/services"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// MeterHandler handles HTTP requests for meter management.
type MeterHandler struct {
	service services.MeterService
}

// NewMeterHandler creates a new instance of MeterHandler.
func NewMeterHandler(service services.MeterService) *MeterHandler {
	return &MeterHandler{
		service: service,
	}
}

// CreateMeter handles POST /meters.
func (h *MeterHandler) CreateMeter(c *gin.Context) {
	var req dto.CreateMeterRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	meter, err := h.service.CreateMeter(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create meter"})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"message": "Meter created successfully", "meter": meter})
}

// ListAllMeters handles GET /meters.
func (h *MeterHandler) ListAllMeters(c *gin.Context) {
	meters, err := h.service.ListAllMeters()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to list meters"})
		return
	}
	c.JSON(http.StatusOK, meters)
}

// GetMeterByID handles GET /meters/:id.
func (h *MeterHandler) GetMeterByID(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid meter ID"})
		return
	}
	meter, err := h.service.GetMeterByID(uint(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve meter"})
		return
	}
	if meter == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Meter not found"})
		return
	}
	c.JSON(http.StatusOK, meter)
}

// UpdateMeter handles PUT /meters/:id.
func (h *MeterHandler) UpdateMeter(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid meter ID"})
		return
	}
	var req dto.UpdateMeterRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := h.service.UpdateMeterByID(uint(id), req); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update meter"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Meter updated successfully"})
}

// DeleteMeter handles DELETE /meters/:id.
func (h *MeterHandler) DeleteMeter(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid meter ID"})
		return
	}
	if err := h.service.DeleteMeterByID(uint(id)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete meter"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Meter deleted successfully"})
}
