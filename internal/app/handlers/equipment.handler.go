package handlers

import (
	"core-api/internal/app/dto"
	"core-api/internal/domain/services"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// EquipmentHandler handles HTTP requests for equipment management.
type EquipmentHandler struct {
	service services.EquipmentService
}

// NewEquipmentHandler creates a new instance of EquipmentHandler.
func NewEquipmentHandler(service services.EquipmentService) *EquipmentHandler {
	return &EquipmentHandler{
		service: service,
	}
}

// CreateEquipment handles POST /equipments.
func (h *EquipmentHandler) CreateEquipment(c *gin.Context) {
	var req dto.CreateEquipmentRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	equip, err := h.service.CreateEquipment(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create equipment"})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"message": "Equipment created successfully", "equipment": equip})
}

// ListAllEquipments handles GET /equipments.
func (h *EquipmentHandler) ListAllEquipments(c *gin.Context) {
	equipments, err := h.service.ListAllEquipments()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to list equipments"})
		return
	}
	c.JSON(http.StatusOK, equipments)
}

// GetEquipmentByID handles GET /equipments/:id.
func (h *EquipmentHandler) GetEquipmentByID(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid equipment ID"})
		return
	}
	equip, err := h.service.GetEquipmentByID(uint(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve equipment"})
		return
	}
	if equip == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Equipment not found"})
		return
	}
	c.JSON(http.StatusOK, equip)
}

// UpdateEquipment handles PUT /equipments/:id.
func (h *EquipmentHandler) UpdateEquipment(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid equipment ID"})
		return
	}
	var req dto.UpdateEquipmentRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := h.service.UpdateEquipmentByID(uint(id), req); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update equipment"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Equipment updated successfully"})
}

// DeleteEquipment handles DELETE /equipments/:id.
func (h *EquipmentHandler) DeleteEquipment(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid equipment ID"})
		return
	}
	if err := h.service.DeleteEquipmentByID(uint(id)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete equipment"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Equipment deleted successfully"})
}
