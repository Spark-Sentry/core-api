package handlers

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"time"

	"core-api/internal/infrastructure/influxdb"
	"github.com/gin-gonic/gin"
)

// WeatherUploadHandler handles the CSV upload of weather data.
type WeatherUploadHandler struct {
	influxClient influxdb.ClientInfluxDBClient
}

// NewWeatherUploadHandler creates a new instance of WeatherUploadHandler.
func NewWeatherUploadHandler(influxClient influxdb.ClientInfluxDBClient) *WeatherUploadHandler {
	return &WeatherUploadHandler{
		influxClient: influxClient,
	}
}

// UploadWeatherCSV handles POST /weather/upload.
// It reads the uploaded CSV file, parses its content, and writes measurements to InfluxDB.
func (h *WeatherUploadHandler) UploadWeatherCSV(c *gin.Context) {
	file, err := c.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "CSV file is required"})
		return
	}

	f, err := file.Open()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to open CSV file"})
		return
	}
	defer f.Close()

	reader := csv.NewReader(bufio.NewReader(f))
	// Skip header line
	if _, err := reader.Read(); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to read CSV header"})
		return
	}

	// Process each CSV record
	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("Error reading CSV: %v", err)})
			return
		}
		// Expected columns (example):
		// 0: longitude, 1: latitude, 2: station_name, 3: climate_id, 4: datetime,
		// 9: temperature, 11: dew_point, 13: humidity, 23: pressure, 29: weather

		fmt.Println(record)
		stationName := record[2]
		climateID := record[3]
		datetime := record[4]
		temperature, _ := strconv.ParseFloat(record[5], 64)
		dewPoint, _ := strconv.ParseFloat(record[6], 64)
		humidity, _ := strconv.ParseFloat(record[7], 64)
		pressure, _ := strconv.ParseFloat(record[8], 64)
		weatherDesc := record[9]

		t, err := time.Parse(time.RFC3339, datetime)
		if err != nil {
			continue
		}

		tags := map[string]string{
			"station_id":   climateID,
			"station_name": stationName,
			"weather":      weatherDesc,
		}
		fields := map[string]interface{}{
			"temperature": temperature,
			"dew_point":   dewPoint,
			"humidity":    humidity,
			"pressure":    pressure,
		}

		if err := h.influxClient.WritePoint("weather_measurements", tags, fields, t); err != nil {
			fmt.Printf("Failed to write point: %v\n", err)
		}
	}

	c.JSON(http.StatusOK, gin.H{"message": "Weather data uploaded successfully"})
}
