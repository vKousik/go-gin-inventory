package usecase

import (
	"context"

	"github.com/vKousik/go-gin-inventory/internal/domain"
	"github.com/vKousik/go-gin-inventory/internal/domain/customErrors"
)

type CategoryUsecase struct {
	repo domain.CategoryRepository
}

func NewCategoryUsecase(repo domain.CategoryRepository) *CategoryUsecase {
	return &CategoryUsecase{repo: repo}
}

func (u *CategoryUsecase) GetAll(ctx context.Context) ([]domain.Category, error) {
	categories, err := u.repo.GetAll(ctx)
	if err != nil {
		return nil, err
	}

	if len(categories) == 0 {
		return nil, customErrors.ErrNotFound
	}
	return u.repo.GetAll(ctx)
}
