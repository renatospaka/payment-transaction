package repository

import (
	"github.com/renatospaka/payment-transaction/core/entity"
)

type TransactionInterface interface {
	Create(transaction *entity.Transaction) error
	Delete(transactionId string) error
	Update(transaction *entity.Transaction) error
	Find(transactionId string) (*entity.Transaction, error)
	FindAll(page, limit int, sort string) ([]*entity.Transaction, error)
}
