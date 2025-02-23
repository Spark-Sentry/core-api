package dto

// CreateBillRequest represents the data required to create a new bill.
type CreateBillRequest struct {
	MeterID  uint    `json:"meterId" binding:"required"`
	Start    string  `json:"start" binding:"required"` // Expected format (e.g., RFC3339)
	Stop     string  `json:"stop" binding:"required"`  // Expected format
	Quantity float64 `json:"quantity" binding:"required"`
	Cost     float64 `json:"cost" binding:"required"`
}

// UpdateBillRequest represents the data required to update an existing bill.
type UpdateBillRequest struct {
	Start    string  `json:"start" binding:"required"`
	Stop     string  `json:"stop" binding:"required"`
	Quantity float64 `json:"quantity" binding:"required"`
	Cost     float64 `json:"cost" binding:"required"`
}
