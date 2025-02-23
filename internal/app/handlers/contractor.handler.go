package handlers

import (
	"core-api/internal/app/dto"
	"core-api/internal/domain/services"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// ContractorHandler handles HTTP requests for contractor management.
type ContractorHandler struct {
	contractorService services.ContractorService
}

// NewContractorHandler creates a new instance of ContractorHandler.
func NewContractorHandler(contractorService services.ContractorService) *ContractorHandler {
	return &ContractorHandler{
		contractorService: contractorService,
	}
}

// CreateContractor handles POST /contractors.
func (h *ContractorHandler) CreateContractor(c *gin.Context) {
	var req dto.CreateContractorRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	contractor, err := h.contractorService.CreateContractor(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create contractor"})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"message": "Contractor created successfully", "contractor": contractor})
}

// ListAllContractors handles GET /contractors.
func (h *ContractorHandler) ListAllContractors(c *gin.Context) {
	contractors, err := h.contractorService.ListAllContractors()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to list contractors"})
		return
	}
	c.JSON(http.StatusOK, contractors)
}

// GetContractorByID handles GET /contractors/:id.
func (h *ContractorHandler) GetContractorByID(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid contractor ID"})
		return
	}

	contractor, err := h.contractorService.GetContractorByID(uint(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get contractor"})
		return
	}
	if contractor == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Contractor not found"})
		return
	}
	c.JSON(http.StatusOK, contractor)
}

// UpdateContractor handles PUT /contractors/:id.
func (h *ContractorHandler) UpdateContractor(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid contractor ID"})
		return
	}

	var req dto.UpdateContractorRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.contractorService.UpdateContractor(uint(id), req); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update contractor"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Contractor updated successfully"})
}

// DeleteContractor handles DELETE /contractors/:id.
func (h *ContractorHandler) DeleteContractor(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid contractor ID"})
		return
	}

	if err := h.contractorService.DeleteContractor(uint(id)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete contractor"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Contractor deleted successfully"})
}
