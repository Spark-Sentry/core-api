package handlers

import (
	"core-api/internal/app/dto"
	"core-api/internal/domain/services"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// WeatherStationHandler handles HTTP requests for weather station operations.
type WeatherStationHandler struct {
	service services.WeatherStationService
}

// NewWeatherStationHandler creates a new instance of WeatherStationHandler.
func NewWeatherStationHandler(service services.WeatherStationService) *WeatherStationHandler {
	return &WeatherStationHandler{
		service: service,
	}
}

// CreateWeatherStation handles POST /weatherstations.
func (h *WeatherStationHandler) CreateWeatherStation(c *gin.Context) {
	var req dto.CreateWeatherStationRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ws, err := h.service.CreateWeatherStation(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create weather station"})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"message": "Weather station created successfully", "weatherStation": ws})
}

// ListWeatherStations handles GET /weatherstations.
func (h *WeatherStationHandler) ListWeatherStations(c *gin.Context) {
	wsList, err := h.service.ListWeatherStations()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to list weather stations"})
		return
	}
	c.JSON(http.StatusOK, wsList)
}

// GetWeatherStationByID handles GET /weatherstations/:id.
func (h *WeatherStationHandler) GetWeatherStationByID(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid weather station ID"})
		return
	}
	ws, err := h.service.GetWeatherStationByID(uint(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve weather station"})
		return
	}
	if ws == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Weather station not found"})
		return
	}
	c.JSON(http.StatusOK, ws)
}

// UpdateWeatherStation handles PUT /weatherstations/:id.
func (h *WeatherStationHandler) UpdateWeatherStation(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid weather station ID"})
		return
	}
	var req dto.UpdateWeatherStationRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := h.service.UpdateWeatherStationByID(uint(id), req); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update weather station"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Weather station updated successfully"})
}

// DeleteWeatherStation handles DELETE /weatherstations/:id.
func (h *WeatherStationHandler) DeleteWeatherStation(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid weather station ID"})
		return
	}
	if err := h.service.DeleteWeatherStationByID(uint(id)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete weather station"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Weather station deleted successfully"})
}
