package repository

import (
	"context"

	"gorm.io/gorm"

	"github.com/vKousik/go-gin-inventory/internal/domain"
	"github.com/vKousik/go-gin-inventory/internal/domain/customErrors"
)

type categoryRepository struct {
	db *gorm.DB
}

func NewCategoryRepository(db *gorm.DB) domain.CategoryRepository {
	return &categoryRepository{db: db}
}

func (r *categoryRepository) GetAll(ctx context.Context) ([]domain.Category, error) {
	var categories []domain.Category

	err := r.db.WithContext(ctx).
		Table("categories").
		Find(&categories).Error

	if err != nil {
		return nil, err
	}

	if len(categories) == 0 {
		return nil, customErrors.ErrNotFound
	}

	return categories, nil
}
