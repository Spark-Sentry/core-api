package dto

// SavingsParams defines the input JSON parameters for the "GetSavings" request.
type SavingsParams struct {
	TimeStart           string   `json:"timeStart" binding:"required"`
	TimeStop            string   `json:"timeStop" binding:"required"`
	IdEfficiencyMeasure string   `json:"idEfficiencyMeasure" binding:"required"`
	Units               []string `json:"units" binding:"required,dive"`
	Mesh                string   `json:"mesh" binding:"required"`
}

// EfficiencyMeasureByMeasurementParams defines the query parameters for the
type EfficiencyMeasureByMeasurementParams struct {
	TimeStart           string   `json:"timeStart"`             // Start time in ISO8601 format
	TimeStop            string   `json:"timeStop"`              // Stop time in ISO8601 format
	EfficiencyMeasureID []string `json:"efficiency_measure_id"` // Array of efficiency measure IDs
	Mesh                string   `json:"mesh"`                  // Aggregation window: "hourly", "daily", "monthly", "annually" or "total"
}

// ProjectByRegressionParams defines the query parameters for the GET /savings/project_by_regression endpoint.
type ProjectByRegressionParams struct {
	TimeStart    string `json:"timeStart"`    // Start time in ISO8601 format
	TimeStop     string `json:"timeStop"`     // Stop time in ISO8601 format
	ProjectId    string `json:"projectId"`    // Project ID as string (convert as needed)
	RegressionId string `json:"regressionId"` // Regression ID as string (convert as needed)
	Mesh         string `json:"mesh"`         // Aggregation window: "monthly", "annually" or "total"
}
