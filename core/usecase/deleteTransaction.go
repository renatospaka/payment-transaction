package usecase

import (
	"context"
	"errors"
)

func (t *TransactionUsecase) deleteTransaction(ctx context.Context, id string) error {
	if id == "" {
		return errors.New("id is required")
	}

	err := t.repo.Delete(id)
	if err != nil {
		return err
	}

	return nil
}
