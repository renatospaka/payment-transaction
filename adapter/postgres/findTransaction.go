package repository

import (
	"context"
	"database/sql"
	"log"

	"github.com/renatospaka/payment-transaction/core/dto"
	"github.com/renatospaka/payment-transaction/core/entity"
)

func (p *PostgresDatabase) findTransaction(ctx context.Context, id string) (*entity.Transaction, error) {
	log.Println("repository.transactions.findTransaction")

	query := `
	SELECT id, status, value , approved_at, denied_at, created_at, updated_at, deleted_at
	FROM transactions
	WHERE id = $1`
	stmt, err := p.DB.PrepareContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	rows := stmt.QueryRow(id)
	if err != nil {
		return nil, err
	}

	var tr dto.TransactionDto
	err = rows.Scan(&tr.ID, &tr.Status, &tr.Value, 
										&tr.ApprovedAt, 
										&tr.DeniedAt, 
										&tr.CreatedAt, 
										&tr.UpdatedAt, 
										&tr.DeletedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, err 	// is this necessary????)
		}
		return nil, err
	}

	transaction, err := entity.MountTransaction(
		tr.ID,
		tr.Status,
		tr.Value,
		tr.DeletedAt,
		tr.ApprovedAt,
		tr.CreatedAt,
		tr.UpdatedAt,
		tr.DeletedAt,
	)
	if err != nil || !transaction.IsValid() {
		return nil, err
	}

	return transaction, nil
}
