package handlers

import (
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
	var account entities.Account
	if err := c.ShouldBindJSON(&account); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.accountService.CreateAccount(&account); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create account"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Account created successfully", "account": account})
}

func (h *AccountHandler) AssociateUserToAccount(c *gin.Context) {
	var request struct {
		UserID    uint `json:"userId"`
		AccountID uint `json:"accountId"`
	}
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
