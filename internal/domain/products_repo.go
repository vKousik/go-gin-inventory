package domain

import "context"

type ProductRepository interface {
	GetAll(ctx context.Context) ([]Product, error)
	Create(ctx context.Context, product *Product) error
}
