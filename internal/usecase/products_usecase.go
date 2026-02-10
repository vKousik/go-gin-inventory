package usecase

import (
	"context"
	"errors"

	"github.com/google/uuid"
	"github.com/vKousik/go-gin-inventory/internal/domain"
	"github.com/vKousik/go-gin-inventory/internal/domain/customErrors"
)

type ProductUsecase struct {
	repo domain.ProductRepository
}

func NewProductUsecase(repo domain.ProductRepository) *ProductUsecase {
	return &ProductUsecase{repo: repo}
}

func (u *ProductUsecase) GetAll(ctx context.Context) ([]domain.Product, error) {
	return u.repo.GetAll(ctx)
}
func (u *ProductUsecase) Create(ctx context.Context, product *domain.Product) error {
	product.ID = uuid.New()

	if product.ProductName == "" {
		return customErrors.ErrInvalidProductName
	}

	if product.Price <= 0 {
		return customErrors.ErrInvalidPrice
	}

	if product.Quantity < 0 {
		return customErrors.ErrInvalidQuantity
	}

	if product.CategoryID == uuid.Nil {
		return customErrors.ErrInvalidCategoryID
	}

	// persist
	if err := u.repo.Create(ctx, product); err != nil {
		if errors.Is(err, customErrors.ErrNotFound) {
			return customErrors.ErrInvalidCategoryID
		}
		return customErrors.ErrInternalServer
	}

	return nil
}
