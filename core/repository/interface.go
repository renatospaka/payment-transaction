package repository

import (
	"github.com/renatospaka/payment-transaction/core/dto"
)

type TransactionInterface interface {
	Create(transaction *dto.TransactionCreateDto) error
	Delete(transactionId *dto.TransactionDeleteDto) error
	Update(transaction *dto.TransactionUpdateDto) error
	Find(transactionId *dto.TransactionFindDto) (*dto.TransactionDto, error)
	FindAll(page, limit int, sort string) ([]*dto.TransactionDto, error)
}
