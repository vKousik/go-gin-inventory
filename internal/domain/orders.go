package domain

import (
	"time"

	"github.com/google/uuid"
)

type Order struct {
	ID        uuid.UUID `json:"id"`
	ProductID uuid.UUID `json:"product_id"`
	Quantity  int       `json:"quantity"`
	OrderedAt time.Time `json:"ordered_at"`
}
