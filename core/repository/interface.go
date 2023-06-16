package repository

import (
	"github.com/renatospaka/payment-transaction/core/entity"
)

type TransactionInterface interface {
	Create(transaction *entity.Transaction) error
	Delete(id string) error
	Update(transaction *entity.Transaction) error
	Find(id string) (*entity.Transaction, error)
	FindAll(page, limit int, sort string) ([]*entity.Transaction, error)
	Approve(transaction *entity.Transaction) error
	Deny(transaction *entity.Transaction) error
}

// type AuthorizationInterface interface {
// 	approveTransaction(transaction *entity.Transaction) error
// 	denyTransaction(transaction *entity.Transaction) error
// }
