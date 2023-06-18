package usecase

import (
	"context"
	"errors"

	utils "github.com/renatospaka/payment-transaction/utils/entity"
	"github.com/renatospaka/payment-transaction/core/dto"
	"github.com/renatospaka/payment-transaction/utils/dateTime"
)

func (t *TransactionUsecase) findTransaction(ctx context.Context, id string) (*dto.TransactionDto, error) {
	if id == "" {
		return nil, errors.New("id is required")
	}

	if _, err := utils.Parse(id); err != nil {
		return nil, errors.New("invalid id")
	}

	transaction, err := t.repo.Find(id)
	if err != nil {
		return nil, err
	} 
	if transaction == nil {
		return nil, errors.New("transaction id not found")
	}

	tr := &dto.TransactionDto{
		ID:         transaction.GetID(),
		Status:     transaction.GetStatus(),
		Value:      transaction.GetValue(),
		DeniedAt:   dateTime.FormatDateToNull(transaction.DeniedAt()),
		ApprovedAt: dateTime.FormatDateToNull(transaction.ApprovedAt()),
		CreatedAt:  dateTime.FormatDateToNull(transaction.CreatedAt()),
		UpdatedAt:  dateTime.FormatDateToNull(transaction.UpdatedAt()),
		DeletedAt:  dateTime.FormatDateToNull(transaction.DeletedAt()),
	}
	return tr, nil
}
