package dto

import "time"

// EquipmentData represents the data for a single equipment
type EquipmentData struct {
	EquipmentID string    `json:"equipment_id"` // ID of the equipment
	Value       float64   `json:"value"`        // value of the equipment data
	Timestamp   time.Time `json:"timestamp"`    // timestamp of the data collection
}

// DailyCollectionRequest represents the request for daily collection of data
type DailyCollectionRequest struct {
	Date       time.Time       `json:"date"`       // date for which data is requested
	Equipments []EquipmentData `json:"equipments"` // list of equipments for which data is requested
}
