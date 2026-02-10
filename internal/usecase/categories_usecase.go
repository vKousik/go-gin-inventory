package usecase

import (
	"context"

	"github.com/vKousik/go-gin-inventory/internal/domain"
)

type CategoryUsecase struct {
	repo domain.CategoryRepository
}

func NewCategoryUsecase(repo domain.CategoryRepository) *CategoryUsecase {
	return &CategoryUsecase{repo: repo}
}

func (u *CategoryUsecase) GetAll(ctx context.Context) ([]domain.Category, error) {
	return u.repo.GetAll(ctx)
}
