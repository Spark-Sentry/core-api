package dto

type CreateAccountRequest struct {
	Name         string `json:"name" binding:"required"`
	ContactEmail string `json:"contactEmail" binding:"required,email"`
	ContactPhone string `json:"contactPhone"`
	Plan         string `json:"plan" binding:"required"`
}

type AssociateUserAccountRequest struct {
	UserID    uint `json:"userId"`
	AccountID uint `json:"accountId"`
}
