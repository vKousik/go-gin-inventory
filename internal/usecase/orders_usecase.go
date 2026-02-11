package usecase

import (
	"context"

	"github.com/vKousik/go-gin-inventory/internal/domain"
	"github.com/vKousik/go-gin-inventory/internal/domain/customErrors"
)

type OrderUsecase struct {
	repo domain.OrderRepository
}

func NewOrderUsecase(repo domain.OrderRepository) *OrderUsecase {
	return &OrderUsecase{repo: repo}
}

func (u *OrderUsecase) GetAll(ctx context.Context) ([]domain.Order, error) {
	orders, err := u.repo.GetAll(ctx)
	if err != nil {
		return nil, err
	}
	if len(orders) == 0 {
		return nil, customErrors.ErrNotFound
	}
	return orders, nil
}
