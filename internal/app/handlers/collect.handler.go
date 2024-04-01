package handlers

import (
	"core-api/internal/app/dto"
	"core-api/internal/domain/services"
	"github.com/gin-gonic/gin"
	"net/http"
)

// CollectHandler handles incoming data collection requests for equipments.
type CollectHandler struct {
	collectService *services.CollectService
}

// NewCollectHandler creates a new instance of CollectHandler.
func NewCollectHandler(collectService *services.CollectService) *CollectHandler {
	return &CollectHandler{collectService: collectService}
}

// HandleCollectData processes the collect data request for a specific parameter with multiple values.
func (h *CollectHandler) HandleCollectData(c *gin.Context) {
	var dataBatch dto.ParameterValuesBatch
	if err := c.ShouldBindJSON(&dataBatch); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	if err := h.collectService.ProcessParameterValuesBatch(dataBatch); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to process data"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Data processed successfully for the parameter"})
}
