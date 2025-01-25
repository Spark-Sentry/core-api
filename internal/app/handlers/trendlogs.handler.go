package handlers

import (
	"context"
	"core-api/internal/app/dto"
	"core-api/internal/domain/services"
	"github.com/gin-gonic/gin"
	"net/http"
)

// TrendlogsHandler defines the handler for trend logs operations.
type TrendlogsHandler struct {
	trendlogsService *services.TrendlogsService
}

// NewTrendlogsHandler creates a new instance of TrendlogsHandler.
func NewTrendlogsHandler(trendlogsService *services.TrendlogsService) *TrendlogsHandler {
	return &TrendlogsHandler{
		trendlogsService: trendlogsService,
	}
}

// GetTrendlogs retrieves trend log data from InfluxDB and returns it to the client.
func (h *TrendlogsHandler) GetTrendlogs(c *gin.Context) {
	var params dto.TrendlogsParams
	if err := c.ShouldBindJSON(&params); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON request body: " + err.Error()})
		return
	}

	// Validate minimal required fields
	if params.Bucket == "" || params.TimeStart == "" || params.TimeStop == "" || len(params.IdParameters) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Missing required fields (bucket, timeStart, timeStop, idParameters)"})
		return
	}

	dataPoints, err := h.trendlogsService.RetrieveTrendlogs(context.Background(), params)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve trend logs: " + err.Error()})
		return
	}

	c.JSON(http.StatusOK, dataPoints)
}
