package repository

import (
	"context"

	"gorm.io/gorm"

	"github.com/vKousik/go-gin-inventory/internal/domain"
)

type taskRepository struct {
	db *gorm.DB
}

func NewTaskRepository(db *gorm.DB) domain.TaskRepository {
	return &taskRepository{db: db}
}

func (r *taskRepository) GetAll(
	ctx context.Context,
) (
	[]domain.Product,
	[]domain.Category,
	[]domain.Order,
	error,
) {

	var (
		products   []domain.Product
		categories []domain.Category
		orders     []domain.Order
	)

	if err := r.db.WithContext(ctx).Table("products").Find(&products).Error; err != nil {
		return nil, nil, nil, err
	}

	if err := r.db.WithContext(ctx).Table("categories").Find(&categories).Error; err != nil {
		return nil, nil, nil, err
	}

	if err := r.db.WithContext(ctx).Table("orders").Find(&orders).Error; err != nil {
		return nil, nil, nil, err
	}

	return products, categories, orders, nil
}
