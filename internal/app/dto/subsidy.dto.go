package dto

// CreateSubsidyRequest represents the data required to create a new subsidy.
type CreateSubsidyRequest struct {
	Organisation        string  `json:"organisation" binding:"required"`
	DateObtained        string  `json:"dateObtained" binding:"required"`
	Amount              float64 `json:"amount" binding:"required"`
	ProjectID           *uint   `json:"projectId"`           // Optional
	EfficiencyMeasureID *uint   `json:"efficiencyMeasureId"` // Optional
}

// UpdateSubsidyRequest represents the data required to update an existing subsidy.
type UpdateSubsidyRequest struct {
	Organisation        string  `json:"organisation" binding:"required"`
	DateObtained        string  `json:"dateObtained" binding:"required"`
	Amount              float64 `json:"amount" binding:"required"`
	ProjectID           *uint   `json:"projectId"`           // Optional
	EfficiencyMeasureID *uint   `json:"efficiencyMeasureId"` // Optional
}
