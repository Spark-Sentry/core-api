package dto

// CreateBmsRequest represents the data required to create a new Bms.
type CreateBmsRequest struct {
	Type      string `json:"type" binding:"required"`      // e.g. "Compass"
	IDPattern string `json:"idPattern" binding:"required"` // e.g. "{hostDevice}.{device}.{log}.{point}"
}

// UpdateBmsRequest represents the data required to update an existing Bms.
type UpdateBmsRequest struct {
	Type      string `json:"type" binding:"required"`
	IDPattern string `json:"idPattern" binding:"required"`
}
