package repository

import (
	"context"

	"gorm.io/gorm"

	"github.com/vKousik/go-gin-inventory/internal/domain"
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
	return categories, nil
}
func (r *categoryRepository) Exists(ctx context.Context, id string) (bool, error) {
	var count int64

	err := r.db.WithContext(ctx).
		Table("categories").
		Where("id = ?", id).
		Count(&count).Error

	if err != nil {
		return false, err
	}

	return count > 0, nil
}
