package dto

// CreateMeterRequest represents the data required to create a new meter.
type CreateMeterRequest struct {
	Supplier   string `json:"supplier" binding:"required"`
	BuildingID uint   `json:"buildingId" binding:"required"`
	Energy     string `json:"energy" binding:"required"`
	Name       string `json:"name" binding:"required"`
	MeterID    string `json:"meterId" binding:"required"`
}

// UpdateMeterRequest represents the data required to update an existing meter.
type UpdateMeterRequest struct {
	Supplier   string `json:"supplier" binding:"required"`
	BuildingID uint   `json:"buildingId" binding:"required"`
	Energy     string `json:"energy" binding:"required"`
	Name       string `json:"name" binding:"required"`
	MeterID    string `json:"meterId" binding:"required"`
}
