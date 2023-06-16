package usecase

import (
	"context"
	"errors"
	"log"

	"github.com/renatospaka/payment-transaction/core/dto"
	// "github.com/renatospaka/payment-transaction/utils"
)

func (t *TransactionUsecase) findTransaction(ctx context.Context, id string) (*dto.TransactionDto, error) {
	if id == "" {
		return nil, errors.New("id is required")
	}

	transaction, err := t.repo.Find(id)
	if err != nil || !transaction.IsValid() {
		return nil, err
	}

	log.Printf("usecase - findTransaction 1 - CreatedAt: %v, UpdatedAt: %v, DeletedAt: %v\n", transaction.CreatedAt(), transaction.UpdatedAt(), transaction.DeletedAt())

	tr := &dto.TransactionDto{
		ID:         transaction.GetID(),
		Status:     transaction.GetStatus(),
		Value:      transaction.GetValue(),
		DeniedAt:   transaction.DeniedAt().String(),
		ApprovedAt: transaction.ApprovedAt().String(),
		CreatedAt:  transaction.CreatedAt().String(),
		UpdatedAt:  transaction.UpdatedAt().String(),
		DeletedAt:  transaction.DeletedAt().String(),
	}
	log.Printf("usecase - findTransaction 2 - CreatedAt: %v, UpdatedAt: %v, DeletedAt: %v\n", transaction.CreatedAt().String(), transaction.UpdatedAt().String(), transaction.DeletedAt().String())
	return tr, nil
}
