package dto

// AreaCreateDTO represents the data structure for a request to create one or more new areas.
type AreaCreateDTO struct {
	Areas []AreaDTO `json:"areas" binding:"required,dive"`
}

// AreaDTO represents the data for an individual area, without the BuildingID.
type AreaDTO struct {
	Name        string `json:"name" binding:"required"`
	Description string `json:"description"`
}

// AreaUpdateDTO represents the data accepted for updating an area.
type AreaUpdateDTO struct {
	Name        string `json:"name" binding:"required"` // Name of the area, required
	Description string `json:"description"`             // Description of the area, optional
}
