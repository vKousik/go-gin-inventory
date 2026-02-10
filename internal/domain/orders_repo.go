package domain

import "context"

type OrderRepository interface {
	GetAll(ctx context.Context) ([]Order, error)
}
