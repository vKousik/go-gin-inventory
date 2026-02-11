package app

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"github.com/vKousik/go-gin-inventory/internal/delivery/http"
	handlers "github.com/vKousik/go-gin-inventory/internal/delivery/http/handler"
	"github.com/vKousik/go-gin-inventory/internal/delivery/http/middleware"
	"github.com/vKousik/go-gin-inventory/internal/repository"
	"github.com/vKousik/go-gin-inventory/internal/usecase"
)

func NewRouter(db *gorm.DB) *gin.Engine {
	// repositories
	productRepo := repository.NewProductRepository(db)
	orderRepo := repository.NewOrderRepository(db)
	categoryRepo := repository.NewCategoryRepository(db)
	taskRepo := repository.NewTaskRepository(db)

	// usecases
	productUC := usecase.NewProductUsecase(productRepo, categoryRepo)
	orderUC := usecase.NewOrderUsecase(orderRepo)
	categoryUC := usecase.NewCategoryUsecase(categoryRepo)
	taskUC := usecase.NewTaskUsecase(taskRepo)

	// handlers
	productHandler := handlers.NewProductHandler(productUC)
	orderHandler := handlers.NewOrderHandler(orderUC)
	categoryHandler := handlers.NewCategoryHandler(categoryUC)
	taskHandler := handlers.NewTaskHandler(taskUC)

	// router
	r := gin.Default()
	r.Use(middleware.CORS())
	http.RegisterRoutes(r, productHandler, orderHandler, categoryHandler, taskHandler)

	return r
}
