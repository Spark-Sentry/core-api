package dto

// CreateRegressionRequest represents the data required to create a new regression.
type CreateRegressionRequest struct {
	ProjectID    uint      `json:"projectId" binding:"required"`
	Name         string    `json:"name" binding:"required"`
	Unit         string    `json:"unit" binding:"required"`
	Coefficients []float64 `json:"coefficients" binding:"required"` // Each value will create a Coefficient entity.
	Meters       []uint    `json:"meters" binding:"required"`       // List of Meter IDs.
}

// UpdateRegressionRequest represents the data required to update an existing regression.
type UpdateRegressionRequest struct {
	Name         string    `json:"name" binding:"required"`
	Unit         string    `json:"unit" binding:"required"`
	Coefficients []float64 `json:"coefficients" binding:"required"`
	Meters       []uint    `json:"meters" binding:"required"`
}
