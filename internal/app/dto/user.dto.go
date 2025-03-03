package dto

type AccountResponse struct {
	ID           uint   `json:"id"`
	Name         string `json:"name"`
	ContactEmail string `json:"contactEmail,omitempty"`
}

type UserResponse struct {
	ID        uint            `json:"id"`
	Email     string          `json:"email"`
	FirstName string          `json:"first_name,omitempty"`
	LastName  string          `json:"last_name,omitempty"`
	Role      string          `json:"role,omitempty"`
	Account   AccountResponse `json:"account,omitempty"`
}
