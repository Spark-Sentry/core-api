package handlers

import (
	"core-api/internal/app/dto"
	"core-api/internal/domain/entities"
	"core-api/internal/domain/services"
	"github.com/gin-gonic/gin"
	"net/http"
)

type AccountHandler struct {
	accountService *services.AccountService
}

func NewAccountHandler(accountService *services.AccountService) *AccountHandler {
	return &AccountHandler{
		accountService: accountService,
	}
}

func (h *AccountHandler) CreateAccount(c *gin.Context) {
	var req dto.CreateAccountRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userRole, exists := c.Get("userRole")
	if !exists || userRole != "superadmin" {
		c.JSON(http.StatusForbidden, gin.H{"error": "Requires superadmin role"})
		return
	}

	account := entities.Account{
		Name:         req.Name,
		ContactEmail: req.ContactEmail,
		ContactPhone: req.ContactPhone,
		Plan:         req.Plan,
	}

	if err := h.accountService.CreateAccount(&account); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create account"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Account created successfully", "account": account})
}

func (h *AccountHandler) AssociateUserToAccount(c *gin.Context) {
	userRole, exists := c.Get("userRole")
	if !exists || userRole != "superadmin" {
		c.JSON(http.StatusForbidden, gin.H{"error": "Requires superadmin role"})
		return
	}

	var request dto.AssociateUserAccountRequest
	if err := c.BindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	err := h.accountService.AssociateUserToAccount(request.UserID, request.AccountID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to associate user with account"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User successfully associated with account"})
}
