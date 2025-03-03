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

// GetTrendlogs retrieves trend log data based on query parameters.
func (h *TrendlogsHandler) GetTrendlogs(c *gin.Context) {
	// Read query parameters
	bucket := c.Query("bucket")
	start := c.Query("start")
	stop := c.Query("stop")
	mesh := c.Query("mesh")
	idParameters := c.QueryArray("idParameters")

	// Validate required parameters
	if bucket == "" || start == "" || stop == "" || mesh == "" || len(idParameters) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Missing required query parameters: bucket, start, stop, mesh, idParameters"})
		return
	}

	params := dto.TrendlogsParams{
		Bucket:       bucket,
		TimeStart:    start,
		TimeStop:     stop,
		IdParameters: idParameters,
		Mesh:         mesh,
	}

	dataPoints, err := h.trendlogsService.RetrieveTrendlogs(context.Background(), params)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve trend logs: " + err.Error()})
		return
	}

	c.JSON(http.StatusOK, dataPoints)
}
