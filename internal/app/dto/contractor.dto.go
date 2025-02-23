package dto

// CreateContractorRequest defines the data required to create a new contractor.
type CreateContractorRequest struct {
	Name  string `json:"name" binding:"required"`
	Phone string `json:"phone"`
	Email string `json:"email"`
}

// UpdateContractorRequest defines the data required to update an existing contractor.
type UpdateContractorRequest struct {
	Name  string `json:"name" binding:"required"`
	Phone string `json:"phone"`
	Email string `json:"email"`
}
