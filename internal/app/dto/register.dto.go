package dto

// RegisterAccountRequest defines the structure for registering a user along with their associated account.
type RegisterAccountRequest struct {
	Email     string      `json:"email" binding:"required,email"`
	Password  string      `json:"password" binding:"required"`
	FirstName string      `json:"firstName" binding:"required"`
	LastName  string      `json:"lastName" binding:"required"`
	Role      string      `json:"role"`
	Account   AccountInfo `json:"account" binding:"required"`
}

// AccountInfo holds the account information necessary for account creation.
type AccountInfo struct {
	Name         string `json:"name" binding:"required"`
	ContactEmail string `json:"contactEmail"`
	ContactPhone string `json:"contactPhone"`
	Plan         string `json:"plan"` // Ex: Basic, Premium
}
