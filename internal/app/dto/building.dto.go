package dto

// CreateBuildingRequest represents the data required to create a new building.
type CreateBuildingRequest struct {
	Name       string `json:"name" binding:"required"`
	Address    string `json:"address" binding:"required"`
	CategoryID uint   `json:"categoryId" binding:"required"`
}

// UpdateBuildingRequest represents the data required to update an existing building.
type UpdateBuildingRequest struct {
	Name       string `json:"name" binding:"required"`
	Address    string `json:"address" binding:"required"`
	CategoryID uint   `json:"categoryId" binding:"required"`
}
