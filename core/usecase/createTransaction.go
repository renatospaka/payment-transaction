package usecase

import (
	"context"

	// "github.com/renatospaka/payment-transaction/core/dto"
	"github.com/renatospaka/payment-transaction/core/entity"
)

func (t *TransactionUsecase) createTransaction(ctx context.Context, transaction *entity.Transaction) error {
	// tr, err := entity.NewTransaction(transaction.Value)
	// if err != nil {
	// 	return err
	// }

	// trDto := dto.TransactionCreateDto{
	// 	ID:         tr.GetID(),
	// 	Value:      tr.GetValue(),
	// 	Status:     tr.GetStatus(),
	// 	DeniedAt:   tr.DeniedAt(),
	// 	ApprovedAt: tr.ApprovedAt(),
	// 	CreatedAt:  tr.TrailDate.CreatedAt(),
	// 	UpdatedAt:  tr.TrailDate.UpdatedAt(),
	// 	DeletedAt:  tr.TrailDate.DeletedAt(),
	// }
	return nil
}
