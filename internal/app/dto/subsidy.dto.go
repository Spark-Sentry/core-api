package dto

// CreateSubsidyRequest represents the data required to create a new subsidy.
type CreateSubsidyRequest struct {
	Type         string  `json:"type" binding:"required,oneof=project efficiencyMeasure"`
	ID           uint    `json:"id" binding:"required"` // ID of the project or efficiency measure
	Organisation string  `json:"organisation" binding:"required"`
	DateObtained string  `json:"dateObtained" binding:"required"` // Expected datetime format (e.g., RFC3339)
	Amount       float64 `json:"amount" binding:"required"`
}

// UpdateSubsidyRequest represents the data required to update an existing subsidy.
type UpdateSubsidyRequest struct {
	Organisation        string  `json:"organisation" binding:"required"`
	DateObtained        string  `json:"dateObtained" binding:"required"`
	Amount              float64 `json:"amount" binding:"required"`
	ProjectID           *uint   `json:"projectId"`           // Optional
	EfficiencyMeasureID *uint   `json:"efficiencyMeasureId"` // Optional
}
