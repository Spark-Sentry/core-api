package dto

// CreateEfficiencyMeasureRequest represents the data needed to create a new efficiency measure.
type CreateEfficiencyMeasureRequest struct {
	ProjectID       uint    `json:"projectId" binding:"required"`
	Name            string  `json:"name" binding:"required"`
	EfficiencyCost  float64 `json:"efficiencyCost" binding:"required"`
	MaintenanceCost float64 `json:"maintenanceCost" binding:"required"`
}

// UpdateEfficiencyMeasureRequest represents the data needed to update an existing efficiency measure.
type UpdateEfficiencyMeasureRequest struct {
	Name            string  `json:"name" binding:"required"`
	EfficiencyCost  float64 `json:"efficiencyCost" binding:"required"`
	MaintenanceCost float64 `json:"maintenanceCost" binding:"required"`
}
