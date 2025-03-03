package handlers

import (
	"core-api/internal/app/dto"
	"core-api/internal/domain/services"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// BillHandler handles HTTP requests for bill management.
type BillHandler struct {
	service services.BillService
}

// NewBillHandler creates a new instance of BillHandler.
func NewBillHandler(service services.BillService) *BillHandler {
	return &BillHandler{
		service: service,
	}
}

// CreateBill handles POST /bills.
func (h *BillHandler) CreateBill(c *gin.Context) {
	var req dto.CreateBillRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	bill, err := h.service.CreateBill(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create bill"})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"message": "Bill created successfully", "bill": bill})
}

// ListBills handles GET /bills.
func (h *BillHandler) ListBills(c *gin.Context) {
	bills, err := h.service.ListBills()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to list bills"})
		return
	}
	c.JSON(http.StatusOK, bills)
}

// GetBillByID handles GET /bills/:id.
func (h *BillHandler) GetBillByID(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid bill ID"})
		return
	}
	bill, err := h.service.GetBillByID(uint(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve bill"})
		return
	}
	if bill == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Bill not found"})
		return
	}
	c.JSON(http.StatusOK, bill)
}

// UpdateBill handles PUT /bills/:id.
func (h *BillHandler) UpdateBill(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid bill ID"})
		return
	}
	var req dto.UpdateBillRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := h.service.UpdateBillByID(uint(id), req); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update bill"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Bill updated successfully"})
}

// DeleteBill handles DELETE /bills/:id.
func (h *BillHandler) DeleteBill(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid bill ID"})
		return
	}
	if err := h.service.DeleteBillByID(uint(id)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete bill"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Bill deleted successfully"})
}
