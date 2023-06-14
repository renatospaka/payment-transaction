package usecase

import (
	"github.com/renatospaka/payment-transaction/core/dto"
	"github.com/renatospaka/payment-transaction/core/entity"
)

func (t *TransactionUsecase) createTransaction(tr *dto.TransactionCreateDto) error {
	transaction, err := entity.NewTransaction(tr.Value)
	if err != nil || !transaction.IsValid() {
		return err
	}

	return t.repo.Create(transaction)
}
