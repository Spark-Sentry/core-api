package dto

// CreateWeatherStationRequest represents the data required to create a new weather station.
type CreateWeatherStationRequest struct {
	ClimateID string  `json:"climateId" binding:"required"`
	Longitude float64 `json:"longitude" binding:"required"`
	Latitude  float64 `json:"latitude" binding:"required"`
	Name      string  `json:"name" binding:"required"`
}

// UpdateWeatherStationRequest represents the data required to update an existing weather station.
type UpdateWeatherStationRequest struct {
	ClimateID string  `json:"climateId" binding:"required"`
	Longitude float64 `json:"longitude" binding:"required"`
	Latitude  float64 `json:"latitude" binding:"required"`
	Name      string  `json:"name" binding:"required"`
}
