package usecase

import (
	"context"

	"github.com/renatospaka/payment-transaction/core/dto"
	"github.com/renatospaka/payment-transaction/core/entity"
)

func (t *TransactionUsecase) createTransaction(ctx context.Context, tr *dto.TransactionCreateDto) error {
	transaction, err := entity.NewTransaction(tr.Value)
	if err != nil || !transaction.IsValid() {
		return err
	}

	err = t.repo.Create(transaction)
	if err != nil {
		return err
	}

	tr.ID = transaction.GetID()
	return nil
}
