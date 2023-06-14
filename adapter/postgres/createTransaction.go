package repository

import (
	"context"
	"log"
	"time"

	"github.com/renatospaka/payment-transaction/core/entity"
)

func (p *PostgresDatabase) createTransaction(ctx context.Context, tr *entity.Transaction) error {
	log.Println("repository.transactions.createTransaction")
	query := `
	INSERT INTO transactions
		(id, status, value, approved_at, denied_at, created_at, updated_at, deleted_at) 
	VALUES
		($1, $2, $3, $4, $5, $6, $7, $8)
	`
	stmt, err := p.DB.PrepareContext(ctx, query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.ExecContext(ctx, tr.GetID(), tr.GetStatus(), tr.GetValue(), 
												tr.ApprovedAt().Format(time.UnixDate), 
												tr.DeniedAt().Format(time.UnixDate), 
												tr.CreatedAt().Format(time.UnixDate), 
												tr.UpdatedAt().Format(time.UnixDate), 
												tr.DeletedAt().Format(time.UnixDate))
	if err != nil {
		return err
	}
	return nil
}
