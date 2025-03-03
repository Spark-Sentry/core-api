package dto

// TargetDTO represents a target to be associated with a project.
type TargetDTO struct {
	Name  string  `json:"name" binding:"required"`
	Value float64 `json:"value" binding:"required"`
	Type  string  `json:"type" binding:"required"`
}

// CreateProjectRequest represents the data required to create a new project.
type CreateProjectRequest struct {
	BuildingID         uint        `json:"buildingId" binding:"required"`
	Name               string      `json:"name" binding:"required"`
	ContractorID       uint        `json:"contractorId" binding:"required"`
	ImplementationDate string      `json:"implementationDate" binding:"required"` // Expected format RFC3339
	EfficiencyCost     float64     `json:"efficiencyCost" binding:"required"`
	MaintenanceCost    float64     `json:"maintenanceCost" binding:"required"`
	Targets            []TargetDTO `json:"targets" binding:"required,dive,required"`
}

// UpdateProjectRequest represents the data required to update an existing project.
type UpdateProjectRequest struct {
	Name               string  `json:"name" binding:"required"`
	ContractorID       uint    `json:"contractorId" binding:"required"`
	ImplementationDate string  `json:"implementationDate" binding:"required"`
	EfficiencyCost     float64 `json:"efficiencyCost" binding:"required"`
	MaintenanceCost    float64 `json:"maintenanceCost" binding:"required"`
}
