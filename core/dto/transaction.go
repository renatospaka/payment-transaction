package dto

import (
	"time"
)

type TransactionDto struct {
	ID         string    `json:"transaction_id"`
	Status     string    `json:"status"`
	Value      float32   `json:"value"`
	DeniedAt   string `json:"denied_at"`
	ApprovedAt string `json:"approved_at"`
	CreatedAt  string `json:"created_at"`
	UpdatedAt  string `json:"updated_at"`
	DeletedAt  string `json:"deleted_at"`
}

type TransactionCreateDto struct {
	ID    string  `json:"transaction_id"`
	Value float32 `json:"value"`
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
