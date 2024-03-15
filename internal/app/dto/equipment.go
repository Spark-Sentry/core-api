package dto

// EquipmentCreateDTO represents the data needed to create new equipment.
type EquipmentCreateDTO struct {
	Tag         string `json:"tag"`
	Description string `json:"description"`
}

// EquipmentUpdateDTO represents the data structure for an equipment update request.
type EquipmentUpdateDTO struct {
	Tag         string `json:"tag" binding:"required"` // Required tag of the equipment
	Description string `json:"description"`            // Optional description of the equipment
}
