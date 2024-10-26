package handlers

import (
	"core-api/internal/app/dto"
	"core-api/internal/domain/services"
	"github.com/gin-gonic/gin"
	"net/http"
)

type CollectHandler struct {
	collectService *services.CollectService
}

// NewCollectHandler creates a new instance of CollectHandler
func NewCollectHandler(collectService *services.CollectService) *CollectHandler {
	return &CollectHandler{collectService: collectService}
}

// CollectHandler handle the data from the BMS and save to influxDB
func (h *CollectHandler) CollectHandler(c *gin.Context) {
	// Parse the JSON request body
	var requestData dto.DailyCollectionRequest
	if err := c.ShouldBindJSON(&requestData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON"})
		return
	}

	/*	// Call the CollectData method with the parsed data
		err := h.collectService.CollectData(requestData)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to collect data"})
			return
		}*/

	c.JSON(http.StatusOK, gin.H{"message": "Data collected successfully"})
}
