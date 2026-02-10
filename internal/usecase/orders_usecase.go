package usecase

import (
	"context"

	"github.com/vKousik/go-gin-inventory/internal/domain"
)

type OrderUsecase struct {
	repo domain.OrderRepository
}

func NewOrderUsecase(repo domain.OrderRepository) *OrderUsecase {
	return &OrderUsecase{repo: repo}
}

func (u *OrderUsecase) GetAll(ctx context.Context) ([]domain.Order, error) {
	return u.repo.GetAll(ctx)
}
