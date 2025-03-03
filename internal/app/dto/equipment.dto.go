package dto

// CreateEquipmentRequest represents the data required to create new equipment.
type CreateEquipmentRequest struct {
	Name string `json:"name" binding:"required"` // Equipment name
	Tag  string `json:"tag" binding:"required"`  // Equipment tag (unique)
}

// UpdateEquipmentRequest represents the data required to update existing equipment.
type UpdateEquipmentRequest struct {
	Name string `json:"name" binding:"required"`
	Tag  string `json:"tag" binding:"required"`
}
