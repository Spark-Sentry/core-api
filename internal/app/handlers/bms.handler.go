package handlers

import (
	"core-api/internal/app/dto"
	"core-api/internal/domain/services"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// BmsHandler handles HTTP requests for Bms operations.
type BmsHandler struct {
	service services.BmsService
}

// NewBmsHandler creates a new instance of BmsHandler.
func NewBmsHandler(service services.BmsService) *BmsHandler {
	return &BmsHandler{
		service: service,
	}
}

// CreateBms handles POST /bms.
func (h *BmsHandler) CreateBms(c *gin.Context) {
	var req dto.CreateBmsRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	bms, err := h.service.CreateBms(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create Bms"})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"message": "Bms created successfully", "bms": bms})
}

// ListBms handles GET /bms.
func (h *BmsHandler) ListBms(c *gin.Context) {
	bmsList, err := h.service.ListBms()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to list Bms"})
		return
	}
	c.JSON(http.StatusOK, bmsList)
}

// GetBmsByID handles GET /bms/:id.
func (h *BmsHandler) GetBmsByID(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Bms ID"})
		return
	}
	bms, err := h.service.GetBmsByID(uint(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve Bms"})
		return
	}
	if bms == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Bms not found"})
		return
	}
	c.JSON(http.StatusOK, bms)
}

// UpdateBms handles PUT /bms/:id.
func (h *BmsHandler) UpdateBms(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Bms ID"})
		return
	}
	var req dto.UpdateBmsRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := h.service.UpdateBmsByID(uint(id), req); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update Bms"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Bms updated successfully"})
}

// DeleteBms handles DELETE /bms/:id.
func (h *BmsHandler) DeleteBms(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Bms ID"})
		return
	}
	if err := h.service.DeleteBmsByID(uint(id)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete Bms"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Bms deleted successfully"})
}
