package usecase

import (
	"context"
	"errors"

	"github.com/renatospaka/payment-transaction/core/dto"
)

func (t *TransactionUsecase) findTransaction(ctx context.Context, id string) (*dto.TransactionDto, error) {
	if id == "" {
		return nil, errors.New("id is required")
	}

	transaction, err := t.repo.Find(id)
	if err != nil || !transaction.IsValid() {
		return nil, err
	}

	tr := &dto.TransactionDto{
		DeniedAt:   transaction.DeniedAt(),
		ApprovedAt: transaction.ApprovedAt(),
		ID:         transaction.GetID(),
		Status:     transaction.GetStatus(),
		Value:      transaction.GetValue(),
	}
	return tr, nil
}
