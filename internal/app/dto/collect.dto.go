package dto

import "time"

// Measurement represents a single value with its timestamp.
type Measurement struct {
	Value     *float64  `json:"value" binding:"required"`     // Measurement value
	Timestamp time.Time `json:"timestamp" binding:"required"` // Measurement timestamp
}

// DailyCollectionRequest represents the request structure for collecting multiple measurements.
type DailyCollectionRequest struct {
	Measurements    []Measurement `json:"measurements" binding:"required,dive,required"` // Array of measurements
	Name            string        `json:"name" binding:"required"`                       // Parameter name
	HostDevice      int           `json:"hostDevice" binding:"required"`                 // Host device ID
	Device          int           `json:"device" binding:"required"`                     // Measurement device ID
	Log             float64       `json:"log,omitempty"`                                 // Log value
	Point           string        `json:"point,omitempty"`                               // Point value
	IdEquipment     int           `json:"id_equipment" binding:"required"`               // Equipment ID
	IdParameter     int           `json:"id_parameter" binding:"required"`               // Parameter ID
	MeasurementName string        `json:"measurement" binding:"required"`                // Measurement name
}
