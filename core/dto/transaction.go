package dto

import (
	"time"
)

type TransactionDto struct {
	DeniedAt   time.Time `json:"denied_at"`
	ApprovedAt time.Time `json:"approved_at"`
	ID         string    `json:"transaction_id"`
	Status     string    `json:"status"`
	Value      float32   `json:"value"`
}

type TransactionCreateDto struct {
	Value      float32   `json:"value"`
}

type TransactionFindDto struct {
	ID string `json:"transaction_id"`
}

type TransactionUpdateDto struct {
	DeniedAt   time.Time `json:"denied_at"`
	ApprovedAt time.Time `json:"approved_at"`
	ID         string    `json:"transaction_id"`
	Status     string    `json:"status"`
	Value      float32   `json:"value"`
}

type TransactionDeleteDto struct {
	ID string `json:"transaction_id"`
}
