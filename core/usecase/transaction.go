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


func (t *TransactionUsecase) Create(tr *dto.TransactionCreateDto) (*dto.TransactionDto, error) {
	log.Println("usecase.transactions.create")

	ctx := context.Background()

	// Create the Transaction
	err := t.createTransaction(ctx, tr)
	if err != nil {
		return nil, err
	}

	// // Call the Authotization service and update status accordingly
	// auth, err := authorization.Process()
	// if err != nil {
	// 	return nil, err
	// } 
	// authorization_id := authorization.authorization_id

	// if auth.Status == entity.TR_APPROVED {
	// 	err = t.approveTransaction(ctx, tr, authorization_id)
	// } else if auth.Status == entity.TR_DENIED {
	// 	err = t.denyTransaction(ctx, tr, authorization_id)
	// }

	// Find the new Transaction the reply to the caller
	transaction, err := t.findTransaction(ctx, tr.ID)
	if err != nil {
		return nil, err
	}
	return transaction, nil
}


func (t *TransactionUsecase) Find(id string) (*dto.TransactionDto, error) {
	log.Println("usecase.transactions.find")
	return t.findTransaction(context.Background(), id)
}


func (t *TransactionUsecase) Delete(id string) error {
	log.Println("usecase.transactions.delete")
	return t.deleteTransaction(context.Background(), id)
}


func (t *TransactionUsecase) Update(id string, tr *dto.TransactionUpdateDto) error {
	// log.Println("usecase.transactions.update")
	// return t.updateTransaction(context.Background(), id, tr)
	panic("implemente me")
}


func (t *TransactionUsecase) FindAll(page, limit int, sort string) ([]*entity.Transaction, error) {
	panic("implemente me")
}
