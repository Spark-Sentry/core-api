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

// GetEfficiencyMeasureByMeasurement handles the GET request for savings by efficiency measure.
func (h *SavingsHandler) GetEfficiencyMeasureByMeasurement(c *gin.Context) {
	// Read query parameters
	start := c.Query("start")
	stop := c.Query("stop")
	mesh := c.Query("mesh")
	measureIDs := c.QueryArray("efficiency_measure_id")

	// Validate required parameters
	if start == "" || stop == "" || mesh == "" || len(measureIDs) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Missing required query parameters: start, stop, mesh, efficiency_measure_id",
		})
		return
	}

	// Build the DTO from query parameters
	params := dto.EfficiencyMeasureByMeasurementParams{
		TimeStart:           start,
		TimeStop:            stop,
		EfficiencyMeasureID: measureIDs,
		Mesh:                mesh,
	}

	// Call the service method with the DTO
	results, err := h.savingsService.RetrieveSavingsForMultipleMeasures(context.Background(), params)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve savings data: " + err.Error()})
		return
	}

	c.JSON(http.StatusOK, results)
}

// GetProjectByRegression handles GET /savings/project_by_regression.
// It computes, for each month between start and stop, the sum over all regression coefficients multiplied by their corresponding independent variable values.
// For "input" source, values are retrieved from the relational DB, and for "computed" source, they are obtained via another service call.
func (h *SavingsHandler) GetProjectByRegression(c *gin.Context) {
	start := c.Query("start")
	stop := c.Query("stop")
	projectId := c.Query("projectId")
	regressionId := c.Query("regressionId")
	mesh := c.Query("mesh")

	if start == "" || stop == "" || projectId == "" || regressionId == "" || mesh == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Missing required query parameters: start, stop, projectId, regressionId, mesh",
		})
		return
	}

	params := dto.ProjectByRegressionParams{
		TimeStart:    start,
		TimeStop:     stop,
		ProjectId:    projectId,
		RegressionId: regressionId,
		Mesh:         mesh,
	}

	results, err := h.savingsService.RetrieveSavingsByRegression(params)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve regression savings: " + err.Error()})
		return
	}

	c.JSON(http.StatusOK, results)
}
