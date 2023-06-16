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

	var (
		tr       dto.TransactionDto
		approved sql.NullTime
		denied   sql.NullTime
		created  sql.NullTime
		updated  sql.NullTime
		deleted  sql.NullTime
	)
	err = rows.Scan(&tr.ID, &tr.Status, &tr.Value,
		&approved,
		&denied,
		&created,
		&updated,
		&deleted,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, err // is this necessary????)
		}
		return nil, err
	}

	log.Printf("postgres - findTransaction 1 - CreatedAt: %v, UpdatedAt: %v, DeletedAt: %v\n", created.Time.String(), updated.Time.String(), deleted.Time.String())
	
	approvedAt := formatNullDate(approved)
	deniedAt := formatNullDate(denied)
	createdAt := formatNullDate(created)
	updatedAt := formatNullDate(updated)
	deletedAt := formatNullDate(deleted)
	log.Printf("postgres - findTransaction 2 - createdAt: %v, updatedAt: %v, deletedAt: %v\n", createdAt, updatedAt, deletedAt)

	transaction, err := entity.MountTransaction(
		tr.ID,
		tr.Status,
		tr.Value,
		deniedAt,
		approvedAt,
		createdAt,
		updatedAt,
		deletedAt,
	)
	if err != nil || !transaction.IsValid() {
		return nil, err
	}

	return transaction, nil
}
