package repository

import (
	"context"
	"log"

	"github.com/renatospaka/payment-transaction/core/entity"
)

func (p *PostgresDatabase) reprocessTransaction(ctx context.Context, tr *entity.Transaction) error {
	log.Println("repository.transactions.reprocessTransaction")

	approvedAt := formatDateToSQL(tr.ApprovedAt())
	deniedAt := formatDateToSQL(tr.DeniedAt())
	createdAt := formatDateToSQL(tr.CreatedAt())
	updatedAt := formatDateToSQL(tr.UpdatedAt())
	deletedAt := formatDateToSQL(tr.DeletedAt())

	query := `
	UPDATE authorizations
	SET		status = $1,
				client_id = $2,
				authorization_id = $3,
				approved_at = $4, 
				denied_at = $5, 
				created_at = $6, 
				updated_at = $7, 
				deleted_at = $8
	WHERE	id = $9
	`
	stmt, err := p.DB.PrepareContext(ctx, query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.ExecContext(ctx,
		tr.GetStatus(),
		tr.GetClientID(),
		tr.GetAuthorizationID(),
		approvedAt,
		deniedAt,
		createdAt,
		updatedAt,
		deletedAt,
		tr.GetID(),
	)
	if err != nil {
		return err
	}
	return nil
}
