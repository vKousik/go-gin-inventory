package usecase

import (
	"context"

	"github.com/vKousik/go-gin-inventory/internal/domain"

	"github.com/vKousik/go-gin-inventory/internal/domain/customErrors"
)

type TaskUsecase struct {
	repo domain.TaskRepository
}

func NewTaskUsecase(repo domain.TaskRepository) *TaskUsecase {
	return &TaskUsecase{repo: repo}
}

func (u *TaskUsecase) GetAll(
	ctx context.Context,
) (
	[]domain.Product,
	[]domain.Category,
	[]domain.Order,
	error,
) {
	products, categories, orders, err := u.repo.GetAll(ctx)
	if err != nil {
		return nil, nil, nil, err
	}
	if len(products) == 0 && len(categories) == 0 && len(orders) == 0 {
		return nil, nil, nil, customErrors.ErrNotFound
	}
	return u.repo.GetAll(ctx)
}
