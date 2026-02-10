package domain

import (
	"time"

	"github.com/google/uuid"
)

type Product struct {
	ID          uuid.UUID `json:"id"`
	ProductName string    `json:"product_name"`
	Price       float64   `json:"price"`
	Quantity    int       `json:"quantity"`
	CategoryID  uuid.UUID `json:"category_id"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"created_at"`
}
