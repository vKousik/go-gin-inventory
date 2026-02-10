package usecase

import (
	"context"

	"github.com/vKousik/go-gin-inventory/internal/domain"
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
	return u.repo.GetAll(ctx)
}
