package handlers

import (
	"core-api/internal/app/dto"
	"core-api/internal/domain/services"
	"github.com/gin-gonic/gin"
	"net/http"
)

type UserHandler struct {
	userService *services.UserService
}

// NewUserHandler creates a new instance of UserHandler.
func NewUserHandler(userService *services.UserService) *UserHandler {
	return &UserHandler{userService: userService}
}

// UsersMe returns the information of the authenticated user and their associated account.
func (h *UserHandler) UsersMe(c *gin.Context) {
	userEmail, exists := c.Get("userEmail") // Retrieving the user's email from the context
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not authenticated"})
		return
	}

	userEmailStr, ok := userEmail.(string)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "User email is not a string"})
		return
	}

	userDetails, err := h.userService.GetUserDetails(userEmailStr)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to retrieve user and account details"})
		return
	}

	response := dto.UserResponse{
		ID:        userDetails.ID,
		Email:     userDetails.Email,
		FirstName: userDetails.FirstName,
		LastName:  userDetails.LastName,
		Role:      userDetails.Role,
		Account: dto.AccountResponse{
			ID:           userDetails.Account.ID,
			Name:         userDetails.Account.Name,
			ContactEmail: userDetails.Account.ContactEmail,
		},
	}

	c.JSON(http.StatusOK, response)
}
