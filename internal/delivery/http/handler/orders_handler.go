package handlers

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/vKousik/go-gin-inventory/internal/domain/customErrors"
	"github.com/vKousik/go-gin-inventory/internal/usecase"
)

type OrderHandler struct {
	uc *usecase.OrderUsecase
}

func NewOrderHandler(uc *usecase.OrderUsecase) *OrderHandler {
	return &OrderHandler{uc: uc}
}

func (h *OrderHandler) GetAll(c *gin.Context) {
	ctx := c.Request.Context()

	orders, err := h.uc.GetAll(ctx)
	if err != nil {

		if errors.Is(err, customErrors.ErrNotFound) {
			c.JSON(http.StatusNotFound, gin.H{
				"message": "no orders found",
				"data":    nil,
			})
			return
		}

		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "internal server error",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "orders fetched successfully",
		"data":    orders,
	})
}
