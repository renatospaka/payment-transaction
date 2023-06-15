package usecase

import (
	"context"
	"log"

	"github.com/renatospaka/payment-transaction/core/dto"
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

func (t *TransactionUsecase) Create(tr *dto.TransactionCreateDto) error {
	log.Println("usecase.transactions.create")
	return t.createTransaction(tr)
}

func (t *TransactionUsecase) Find(id string) (*dto.TransactionDto, error) {
	log.Println("usecase.transactions.find")
	return t.findTransaction(context.Background(), id)
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
