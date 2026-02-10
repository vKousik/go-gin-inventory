package handlers

import (
	"net/http"

	"errors"

	"github.com/gin-gonic/gin"
	"github.com/vKousik/go-gin-inventory/internal/domain/customErrors"
	"github.com/vKousik/go-gin-inventory/internal/usecase"
)

type CategoryHandler struct {
	uc *usecase.CategoryUsecase
}

func NewCategoryHandler(uc *usecase.CategoryUsecase) *CategoryHandler {
	return &CategoryHandler{uc: uc}
}

func (h *CategoryHandler) GetAll(c *gin.Context) {
	ctx := c.Request.Context()

	categories, err := h.uc.GetAll(ctx)
	if err != nil {

		if errors.Is(err, customErrors.ErrNotFound) {
			c.JSON(http.StatusNotFound, gin.H{
				"message": "no categories found",
			})
			return
		}

		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "internal server error",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "categories fetched successfully",
		"data":    categories,
	})
}
