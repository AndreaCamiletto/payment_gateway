package model

import (
	"time"

	"github.com/google/uuid"
)

type Payment struct {
	ID        uuid.UUID    `json:"id"`
	Amount    int64        `json:"amount"`
	Currency  string       `json:"currency"`
	CreatedAt time.Time    `json:"createdAt"`
	UpdatedAt time.Time    `json:"updatedAt"`
	Status    PaymentState `json:"status"`
}
