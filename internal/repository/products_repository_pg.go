package repository

import (
	"context"

	"gorm.io/gorm"

	"github.com/vKousik/go-gin-inventory/internal/domain"
	"github.com/vKousik/go-gin-inventory/internal/domain/customErrors"
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
	if len(products) == 0 {
		return nil, customErrors.ErrNotFound
	}

	return products, nil
}
func (r *productRepository) Create(ctx context.Context, product *domain.Product) error {
	var count int64
	if err := r.db.WithContext(ctx).
		Table("categories").
		Where("id = ?", product.CategoryID).
		Count(&count).Error; err != nil {
		return err
	}
	if count == 0 {
		return customErrors.ErrNotFound
	}
	return r.db.WithContext(ctx).
		Table("products").
		Create(product).
		Error
}
