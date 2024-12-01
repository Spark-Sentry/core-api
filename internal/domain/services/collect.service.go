package services

import (
	"core-api/internal/app/dto"
	"core-api/internal/infrastructure/influxdb"
	"fmt"
	"time"
)

// CollectService provides data collection functionalities.
type CollectService struct {
	influxClient influxdb.ClientInfluxDBClient
}

// NewCollectService initializes a new CollectService with an InfluxDB client.
func NewCollectService(client influxdb.ClientInfluxDBClient) *CollectService {
	return &CollectService{influxClient: client}
}

// CollectData processes each value in request data and saves it to InfluxDB with the provided timestamp.
func (s *CollectService) CollectData(idParam string, data dto.DailyCollectionRequest, value float64, timestamp time.Time) error {
	// Prepare tags and fields for the data point
	tags := map[string]string{
		"id_param":     idParam,
		"name":         data.Name,
		"measurement":  data.MeasurementName,
		"host_device":  fmt.Sprintf("%d", data.HostDevice),
		"device":       fmt.Sprintf("%d", data.Device),
		"id_equipment": fmt.Sprintf("%d", data.IdEquipment),
		"id_parameter": fmt.Sprintf("%d", data.IdParameter),
	}
	fields := map[string]interface{}{
		"value": value, // Individual value from Measurements array
		"log":   data.Log,
		"point": data.Point,
	}

	// Write data point to InfluxDB with specific timestamp
	return s.influxClient.WritePoint(data.MeasurementName, tags, fields, timestamp)
}
