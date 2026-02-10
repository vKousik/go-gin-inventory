package customErrors

import "errors"

var (
	ErrInvalidProductName = errors.New("product name is required")
	ErrInvalidPrice       = errors.New("price must be greater than 0")
	ErrInvalidQuantity    = errors.New("quantity must be non-negative")
	ErrInvalidCategoryID  = errors.New("invalid category_id")
)

var ErrNotFound = errors.New("product not found")

var (
	ErrInternalServer = errors.New("internal server error")
)
