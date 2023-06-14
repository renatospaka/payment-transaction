package usecase

import (
	"context"

	"github.com/renatospaka/payment-transaction/core/entity"
	"github.com/renatospaka/payment-transaction/core/repository"
)

type TransactionUsecase struct {
	repo repository.TransactionInterface
}

func NewTransactionUsecase(repo repository.TransactionInterface) *TransactionUsecase {
	return &TransactionUsecase{
		repo: repo,
	}
}

func (t *TransactionUsecase) Create(transaction *entity.Transaction) error {
	return t.createTransaction(context.Background(), transaction)
}

func (t *TransactionUsecase) Find(transactionId string) (*entity.Transaction, error) {
	return t.findTransaction(context.Background(), transactionId)
}

func (t *TransactionUsecase) Delete(transactionId string) error {
	panic("implemente me")
}

func (t *TransactionUsecase) Update(transaction *entity.Transaction) error {
	panic("implemente me")
}

func (t *TransactionUsecase) FindAll(page, limit int, sort string) ([]*entity.Transaction, error) {
	panic("implemente me")
}
