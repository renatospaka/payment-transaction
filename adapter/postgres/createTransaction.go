package repository

import (
	"context"

	"github.com/renatospaka/payment-transaction/core/dto"
)

func (p *PostgresDatabase) createTransaction(transaction *dto.TransactionCreateDto) error {
	ctx := context.Background()
	query := `
	INSERT INTO public.transactions (
		id, status, value, approved_at, denied_at, created_at, updated_at, deleted_at
	) VALUES (
		?, ?, ?, ?, ?, ?, ?, ?
	)`
	stmt, err := p.DB.PrepareContext(ctx, query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.ExecContext(ctx, transaction.ID, transaction.Status, transaction.Value, transaction.ApprovedAt, transaction.DeniedAt, transaction.CreatedAt, transaction.UpdatedAt, transaction.DeletedAt)
	if err != nil {
		return err
	}
	return nil
}
