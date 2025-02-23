package dto

// CreateIndependantVariableRequest represents the data needed to create a new independant variable.
type CreateIndependantVariableRequest struct {
	Name      string  `json:"name" binding:"required"`
	DateStart string  `json:"dateStart" binding:"required"`
	DateStop  string  `json:"dateStop" binding:"required"`
	Value     float64 `json:"value" binding:"required"`
	Source    string  `json:"source" binding:"required"` // "input" or "computed"
}

// UpdateIndependantVariableRequest represents the data needed to update an existing independant variable.
type UpdateIndependantVariableRequest struct {
	Name      string  `json:"name" binding:"required"`
	DateStart string  `json:"dateStart" binding:"required"`
	DateStop  string  `json:"dateStop" binding:"required"`
	Value     float64 `json:"value" binding:"required"`
	Source    string  `json:"source" binding:"required"`
}
