package http

import (
	"github.com/gin-gonic/gin"
	handlers "github.com/vKousik/go-gin-inventory/internal/delivery/http/handler"
)

func RegisterRoutes(
	r *gin.Engine,
	productHandler *handlers.ProductHandler,
	orderHandler *handlers.OrderHandler,
	categoryHandler *handlers.CategoryHandler,
	taskHandler *handlers.TaskHandler,
) {
	// health check (optional but nice)
	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "ok"})
	})

	// core endpoints
	r.GET("/products", productHandler.GetAll)
	r.GET("/orders", orderHandler.GetAll)
	r.GET("/categories", categoryHandler.GetAll)
	r.GET("/tasks", taskHandler.GetAll)
	r.POST("/products", productHandler.Create)

}
