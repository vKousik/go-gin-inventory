package handlers

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/vKousik/go-gin-inventory/internal/domain/customErrors"
	"github.com/vKousik/go-gin-inventory/internal/usecase"
)

type TaskHandler struct {
	uc *usecase.TaskUsecase
}

func NewTaskHandler(uc *usecase.TaskUsecase) *TaskHandler {
	return &TaskHandler{uc: uc}
}

func (h *TaskHandler) GetAll(c *gin.Context) {
	ctx := c.Request.Context()

	products, categories, orders, err := h.uc.GetAll(ctx)
	if err != nil {

		if errors.Is(err, customErrors.ErrNotFound) {
			c.JSON(http.StatusNotFound, gin.H{
				"message": "no tasks found",
			})
			return
		}

		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "internal server error",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "all data fetched successfully",
		"data": gin.H{
			"products":   products,
			"categories": categories,
			"orders":     orders,
		},
	})
}
