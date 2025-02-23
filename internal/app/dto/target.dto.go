package dto

// CreateTargetRequest represents the data required to create a new target.
type CreateTargetRequest struct {
	Name                string  `json:"name" binding:"required"`
	Value               float64 `json:"value" binding:"required"`
	Type                string  `json:"type" binding:"required"`
	ProjectID           *uint   `json:"projectId"`           // Optional
	EfficiencyMeasureID *uint   `json:"efficiencyMeasureId"` // Optional
}

// UpdateTargetRequest represents the data required to update an existing target.
type UpdateTargetRequest struct {
	Name                string  `json:"name" binding:"required"`
	Value               float64 `json:"value" binding:"required"`
	Type                string  `json:"type" binding:"required"`
	ProjectID           *uint   `json:"projectId"`           // Optional
	EfficiencyMeasureID *uint   `json:"efficiencyMeasureId"` // Optional
}
