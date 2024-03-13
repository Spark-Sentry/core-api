package dto

// CreateBuildingRequest represents the required data to create a Building with at least one Area.
type CreateBuildingRequest struct {
	Name    string              `json:"name"`
	Address string              `json:"address"`
	Group   string              `json:"group"`
	Areas   []CreateAreaRequest `json:"areas"` // Ensure there's at least one Area
}

// CreateAreaRequest represents the data for creating an Area within a Building.
type CreateAreaRequest struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}
