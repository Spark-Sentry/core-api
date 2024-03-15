package dto

// SystemUpdateDTO represents the data structure for a system update request.
type SystemUpdateDTO struct {
	Name        string `json:"name" binding:"required"` // Required name of the system
	Description string `json:"description"`             // Optional description of the system
}
