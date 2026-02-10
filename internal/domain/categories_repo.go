package domain

import "context"

type CategoryRepository interface {
	GetAll(ctx context.Context) ([]Category, error)
}
