package repository

import (
	"context"
	"log"

	"github.com/renatospaka/payment-transaction/core/entity"
)

func (p *PostgresDatabase) createTransaction(ctx context.Context, tr *entity.Transaction) error {
	log.Println("repository.transactions.createTransaction")

	approvedAt := formatDateToSQL(tr.ApprovedAt())
	deniedAt := formatDateToSQL(tr.DeniedAt())
	createdAt := formatDateToSQL(tr.CreatedAt())
	updatedAt := formatDateToSQL(tr.UpdatedAt())
	deletedAt := formatDateToSQL(tr.DeletedAt())

	query := `
	INSERT INTO transactions
		(id, client_id, status, value, approved_at, denied_at, created_at, updated_at, deleted_at) 
	VALUES
		($1, $2, $3, $4, $5, $6, $7, $8, $9)
	`
	stmt, err := p.DB.PrepareContext(ctx, query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.ExecContext(ctx,
		tr.GetID(),
		tr.GetClientID(),
		tr.GetStatus(),
		tr.GetValue(),
		approvedAt,
		deniedAt,
		createdAt,
		updatedAt,
		deletedAt,
	)
	if err != nil {
		return err
	}
	return nil
}
