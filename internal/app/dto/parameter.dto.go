package dto

// CreateParameterRequest represents the data required to create a new parameter.
type CreateParameterRequest struct {
	Name                 string `json:"name" binding:"required"`
	EquipmentID          *uint  `json:"equipmentId"` // Optional
	BmsID                uint   `json:"bmsId" binding:"required"`
	IDInBms              string `json:"idInBms" binding:"required"`
	PointType            string `json:"pointType" binding:"required"`
	Unit                 string `json:"unit" binding:"required"`
	EfficiencyMeasureIDs []uint `json:"efficiencyMeasureIds" binding:"required,dive,required"`
}

// UpdateParameterRequest represents the data required to update an existing parameter.
type UpdateParameterRequest struct {
	Name                 string `json:"name" binding:"required"`
	EquipmentID          *uint  `json:"equipmentId"`
	BmsID                uint   `json:"bmsId" binding:"required"`
	IDInBms              string `json:"idInBms" binding:"required"`
	PointType            string `json:"pointType" binding:"required"`
	Unit                 string `json:"unit" binding:"required"`
	EfficiencyMeasureIDs []uint `json:"efficiencyMeasureIds" binding:"required,dive,required"`
}
