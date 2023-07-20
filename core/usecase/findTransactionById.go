package usecase

import (
	"context"
	"log"

	"github.com/renatospaka/payment-transaction/core/dto"
	"github.com/renatospaka/payment-transaction/core/entity"
	"github.com/renatospaka/payment-transaction/utils/dateTime"
	utils "github.com/renatospaka/payment-transaction/utils/entity"
)

func (t *TransactionUsecase) findTransactionById(ctx context.Context, id string) (*dto.TransactionDto, error) {
	log.Println("usecase.transactions.findTransactionById")


	log.Println("usecase.transactions.findTransactionById - GOT HERE")
	if id == "" {
		return nil, entity.ErrTransactionIDIsRequired
	}

	if _, err := utils.Parse(id); err != nil {
		return nil, entity.ErrInvalidTransactionID
	}

	transaction, err := t.repo.Find(id)
	if err != nil {
		return nil, err
	}
	if transaction == nil {
		return nil, entity.ErrTransactionIDNotFound
	}

	tr := &dto.TransactionDto{
		ID:              transaction.GetID(),
		ClientID:        transaction.GetClientID(),
		AuthorizationID: transaction.GetAuthorizationID(),
		Status:          transaction.GetStatus(),
		Value:           transaction.GetValue(),
		DeniedAt:        dateTime.FormatDateToNull(transaction.DeniedAt()),
		ApprovedAt:      dateTime.FormatDateToNull(transaction.ApprovedAt()),
		CreatedAt:       dateTime.FormatDateToNull(transaction.CreatedAt()),
		UpdatedAt:       dateTime.FormatDateToNull(transaction.UpdatedAt()),
		DeletedAt:       dateTime.FormatDateToNull(transaction.DeletedAt()),
	}
	return tr, nil
}
