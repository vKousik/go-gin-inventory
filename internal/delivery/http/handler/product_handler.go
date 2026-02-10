package handlers

import (
	"errors"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/vKousik/go-gin-inventory/internal/domain"
	"github.com/vKousik/go-gin-inventory/internal/domain/customErrors"
	"github.com/vKousik/go-gin-inventory/internal/usecase"
	"gorm.io/gorm"
)

type ProductHandler struct {
	uc *usecase.ProductUsecase
}

func NewProductHandler(uc *usecase.ProductUsecase) *ProductHandler {
	return &ProductHandler{uc: uc}
}

func (h *ProductHandler) GetAll(c *gin.Context) {
	ctx := c.Request.Context()

	products, err := h.uc.GetAll(ctx)
	if err != nil {

		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, gin.H{
				"message": "no products found",
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
		"message": "products fetched successfully",
		"data":    products,
	})
}
func (h *ProductHandler) Create(c *gin.Context) {
	ctx := c.Request.Context()

	var product domain.Product
	if err := c.ShouldBindJSON(&product); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "invalid request body",
		})
		return
	}

	if err := h.uc.Create(ctx, &product); err != nil {

		switch err {

		case customErrors.ErrInvalidProductName,
			customErrors.ErrInvalidPrice,
			customErrors.ErrInvalidQuantity,
			customErrors.ErrInvalidCategoryID:
			c.JSON(http.StatusBadRequest, gin.H{
				"message": err.Error(),
			})
			return

		case customErrors.ErrNotFound:
			c.JSON(http.StatusNotFound, gin.H{
				"message": err.Error(),
			})
			return
		default:
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": "failed to create product",
			})
			return
		}
	}
	// Concurrency: Simulating sending a notification email using a goroutine
	// after a task/product is created, without blocking the HTTP response.
	// This is a dummy implementation to demonstrate Go's concurrency feature.

	go func(pName string) {
		time.Sleep(4 * time.Second)
		log.Println("Notification email sent for product:", pName)
	}(product.ProductName)

	c.JSON(http.StatusCreated, gin.H{
		"message": "product created successfully",
		"data":    product,
	})
}
