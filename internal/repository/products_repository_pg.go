package repository

import (
	"context"

	"gorm.io/gorm"

	"github.com/vKousik/go-gin-inventory/internal/domain"
)

type productRepository struct {
	db *gorm.DB
}

func NewProductRepository(db *gorm.DB) domain.ProductRepository {
	return &productRepository{db: db}
}

func (r *productRepository) GetAll(ctx context.Context) ([]domain.Product, error) {
	var products []domain.Product

	err := r.db.WithContext(ctx).
		Table("products").
		Find(&products).Error

	if err != nil {
		return nil, err
	}
	return products, nil
}
func (r *productRepository) Create(ctx context.Context, product *domain.Product) error {

	return r.db.WithContext(ctx).
		Table("products").
		Create(product).
		Error
}
