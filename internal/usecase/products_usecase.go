package usecase

import (
	"context"

	"github.com/google/uuid"
	"github.com/vKousik/go-gin-inventory/internal/domain"
	"github.com/vKousik/go-gin-inventory/internal/domain/customErrors"
)

type ProductUsecase struct {
	repo         domain.ProductRepository
	categoryRepo domain.CategoryRepository
}

func NewProductUsecase(
	repo domain.ProductRepository,
	categoryRepo domain.CategoryRepository,
) *ProductUsecase {
	return &ProductUsecase{
		repo:         repo,
		categoryRepo: categoryRepo,
	}
}

func (u *ProductUsecase) GetAll(ctx context.Context) ([]domain.Product, error) {
	products, err := u.repo.GetAll(ctx)
	if err != nil {
		return nil, err
	}
	if len(products) == 0 {
		return nil, customErrors.ErrNotFound
	}
	return products, nil
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

	exists, err := u.categoryRepo.Exists(ctx, product.CategoryID.String())
	if err != nil {
		return customErrors.ErrInternalServer
	}

	if !exists {
		return customErrors.ErrInvalidCategoryID
	}

	if err := u.repo.Create(ctx, product); err != nil {
		return customErrors.ErrInternalServer
	}

	return nil
}
