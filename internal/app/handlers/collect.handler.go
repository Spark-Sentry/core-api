package handlers

import (
	"core-api/internal/app/dto"
	"core-api/internal/domain/services"
	"github.com/gin-gonic/gin"
	"net/http"
)

// CollectHandler handles incoming data collection requests.
type CollectHandler struct {
	collectService *services.CollectService
}

// NewCollectHandler creates a new instance of CollectHandler.
func NewCollectHandler(collectService *services.CollectService) *CollectHandler {
	return &CollectHandler{collectService: collectService}
}

// HandleCollectData processes the collect data request.
func (h *CollectHandler) HandleCollectData(c *gin.Context) {
	var collectData dto.CollectData
	if err := c.ShouldBindJSON(&collectData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	err := h.collectService.ProcessData(collectData)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to process data"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Data processed successfully"})
}
