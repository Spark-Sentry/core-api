package dto

// EquipmentCreateDTO represents the data needed to create new equipment.
type EquipmentCreateDTO struct {
	Tag         string `json:"tag"`
	Description string `json:"description"`
}
