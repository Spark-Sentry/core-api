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

	c.JSON(http.StatusOK, gin.H{"token": token})
}

// Register handle register logic
func (h *AuthHandler) Register(c *gin.Context) {
	var req dto.RegisterRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user := entities.User{
		Email:     req.Email,
		Password:  req.Password,
		FirstName: req.FirstName,
		LastName:  req.LastName,
	}

	if err := h.authService.RegisterUser(&user); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not create user"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "User registered successfully"})
}
