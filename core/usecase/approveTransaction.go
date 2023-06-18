package usecase

import (
	"context"
	// "time"

	"github.com/renatospaka/payment-transaction/core/dto"
	// "github.com/renatospaka/payment-transaction/core/entity"
)

func (t *TransactionUsecase) approveTransaction(ctx context.Context, tr *dto.TransactionCreateDto) error {
	// null := time.Time{}
	// now := time.Now()

	// transaction, err := entity.MountTransaction(tr.ID, entity.TR_PENDING, tr.Value, null, now, now, now, null)
	// if err != nil || !transaction.IsValid() {
	// 	return err
	// }
	// transaction.SetStatusToApproved()

	// return t.repo.Approve(transaction)
	return nil
}
