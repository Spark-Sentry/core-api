package handlers

import (
	"core-api/internal/app/dto"
	"core-api/internal/domain/services"
	"fmt"
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

// CollectHandler handles data from the BMS and saves it to InfluxDB.
func (h *CollectHandler) CollectHandler(c *gin.Context) {
	// Parse the JSON request body
	var requestData dto.DailyCollectionRequest
	if err := c.ShouldBindJSON(&requestData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON"})
		return
	}

	// Loop over each measurement in Measurements and send to the service
	for _, measurement := range requestData.Measurements {
		if err := h.collectService.CollectData(requestData, *measurement.Value, measurement.Timestamp); err != nil {
			fmt.Println(err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to collect data"})
			return
		}
	}

	c.JSON(http.StatusOK, gin.H{"message": "Data collected and saved successfully"})
}
