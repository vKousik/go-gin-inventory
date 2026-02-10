package domain

import "context"

type TaskRepository interface {
	GetAll(
		ctx context.Context,
	) (
		products []Product,
		categories []Category,
		orders []Order,
		err error,
	)
}
