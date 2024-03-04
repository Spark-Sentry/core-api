// /internal/interfaces/handlers/auth_handler.go

package handlers

import (
	"core-api/internal/domain/services"
	"github.com/gin-gonic/gin"
	"net/http"
)

// AuthHandler struct qui inclut le service d'authentification
type AuthHandler struct {
	authService *services.AuthService
}

// NewAuthHandler crée une nouvelle instance d'AuthHandler
func NewAuthHandler(authService *services.AuthService) *AuthHandler {
	return &AuthHandler{authService: authService}
}

// Login gère la requête de login
func (h *AuthHandler) Login(c *gin.Context) {
	var loginDetails services.LoginRequest
	if err := c.ShouldBindJSON(&loginDetails); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request data"})
		return
	}

	token, err := h.authService.Authenticate(loginDetails.Username, loginDetails.Password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": token})
}
