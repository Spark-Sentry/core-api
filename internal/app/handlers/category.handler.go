package handlers

import (
	"core-api/internal/domain/services"
	"github.com/gin-gonic/gin"
	"net/http"
)

// CategoryHandler handles HTTP requests for category management.
type CategoryHandler struct {
	service services.CategoryService
}

// NewCategoryHandler creates a new instance of CategoryHandler.
func NewCategoryHandler(service services.CategoryService) *CategoryHandler {
	return &CategoryHandler{
		service: service,
	}
}

// ListCategories handles GET /categories.
func (h *CategoryHandler) ListCategories(c *gin.Context) {
	categories, err := h.service.ListCategories()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to list categories"})
		return
	}
	c.JSON(http.StatusOK, categories)
}
