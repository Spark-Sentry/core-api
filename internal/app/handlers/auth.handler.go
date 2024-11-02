package handlers

import (
	"core-api/internal/app/dto"
	"core-api/internal/domain/entities"
	"core-api/internal/domain/services"
	"github.com/gin-gonic/gin"
	"net/http"
)

// AuthHandler struct for AuthHandler
type AuthHandler struct {
	authService *services.AuthService
}

// NewAuthHandler create a new instance of AuthHandler
func NewAuthHandler(authService *services.AuthService) *AuthHandler {
	return &AuthHandler{authService: authService}
}

// Login handle login logic
func (h *AuthHandler) Login(c *gin.Context) {
	var loginDetails services.LoginRequest
	if err := c.ShouldBindJSON(&loginDetails); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request data"})
		return
	}

	token, err := h.authService.Authenticate(loginDetails.Email, loginDetails.Password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": token})
}

// Register handles the registration of a new user along with their associated account.
func (h *AuthHandler) Register(c *gin.Context) {
	var req dto.RegisterAccountRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Create an account entity using the provided account information
	account := entities.Account{
		Name:         req.Account.Name,
		ContactEmail: req.Account.ContactEmail,
		ContactPhone: req.Account.ContactPhone,
		Plan:         req.Account.Plan,
	}

	// Create a user entity linked to the account
	user := entities.User{
		Email:     req.Email,
		Password:  req.Password,
		FirstName: req.FirstName,
		LastName:  req.LastName,
		Role:      req.Role,
		Account:   &account, // Associate the user with the account
	}

	// Register the user along with the account in the service layer
	if err := h.authService.RegisterUserAndAccount(&user, &account); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not create user and account"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "User and account registered successfully"})
}
