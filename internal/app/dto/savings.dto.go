package dto

// SavingsParams defines the input JSON parameters for the "GetSavings" request.
type SavingsParams struct {
	TimeStart           string   `json:"timeStart" binding:"required"`
	TimeStop            string   `json:"timeStop" binding:"required"`
	IdEfficiencyMeasure string   `json:"idEfficiencyMeasure" binding:"required"`
	Units               []string `json:"units" binding:"required,dive"`
	Mesh                string   `json:"mesh" binding:"required"`
}
