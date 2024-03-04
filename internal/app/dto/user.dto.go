package dto

type AccountResponse struct {
	ID           uint   `json:"id"`
	Name         string `json:"name"`
	ContactEmail string `json:"contactEmail,omitempty"`
}

type UserResponse struct {
	ID        uint            `json:"id"`
	Email     string          `json:"email"`
	FirstName string          `json:"firstName,omitempty"`
	LastName  string          `json:"lastName,omitempty"`
	Role      string          `json:"role,omitempty"`
	Account   AccountResponse `json:"account,omitempty"`
}
