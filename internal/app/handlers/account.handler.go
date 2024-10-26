package handlers

import (
	"core-api/internal/app/dto"
	"core-api/internal/domain/entities"
	"core-api/internal/domain/services"
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type AccountHandler struct {
	accountService *services.AccountService
}

func NewAccountHandler(accountService *services.AccountService) *AccountHandler {
	return &AccountHandler{
		accountService: accountService,
	}
}

// CreateAccount handles the creation of a new account.
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

// ListAllAccounts retrieves a list of all accounts.
func (h *AccountHandler) ListAllAccounts(c *gin.Context) {
	accounts, err := h.accountService.GetAllAccounts()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve accounts"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": accounts})
}

// GetAccountByID retrieves an account by its ID.
func (h *AccountHandler) GetAccountByID(c *gin.Context) {
	idParam := c.Param("id")
	accountID, err := strconv.ParseUint(idParam, 10, 32)
	account, err := h.accountService.GetAccountByID(uint(accountID))
	if err != nil {
		if errors.Is(err, services.ErrAccountNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": "Account not found"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve account"})
		}
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": account})
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
