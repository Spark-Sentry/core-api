package handlers

import (
	"context"
	"core-api/internal/app/dto"
	"core-api/internal/domain/services"
	"github.com/gin-gonic/gin"
	"net/http"
)

// SavingsHandler manages the HTTP requests for retrieving savings data.
type SavingsHandler struct {
	savingsService *services.SavingsService
}

// NewSavingsHandler creates and returns a new instance of SavingsHandler.
func NewSavingsHandler(savingsService *services.SavingsService) *SavingsHandler {
	return &SavingsHandler{savingsService: savingsService}
}

// GetSavings handles the HTTP request to retrieve savings data from InfluxDB.
func (h *SavingsHandler) GetSavings(c *gin.Context) {
	var params dto.SavingsParams

	// Bind the JSON body to our SavingsParams struct
	if err := c.ShouldBindJSON(&params); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON input: " + err.Error()})
		return
	}

	// Retrieve data using the service
	results, err := h.savingsService.RetrieveSavings(context.Background(), params)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve savings data: " + err.Error()})
		return
	}

	// Return the results in JSON format
	c.JSON(http.StatusOK, results)
}
