package repository

import (
	"context"

	"gorm.io/gorm"

	"github.com/vKousik/go-gin-inventory/internal/domain"
)

type orderRepository struct {
	db *gorm.DB
}

func NewOrderRepository(db *gorm.DB) domain.OrderRepository {
	return &orderRepository{db: db}
}

func (r *orderRepository) GetAll(ctx context.Context) ([]domain.Order, error) {
	var orders []domain.Order

	err := r.db.WithContext(ctx).
		Table("orders").
		Find(&orders).Error

	if err != nil {
		return nil, err
	}

	return orders, nil
}
