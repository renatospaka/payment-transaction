package repository

import (
	"context"
	"log"

	"github.com/renatospaka/payment-transaction/core/dto"
	"github.com/renatospaka/payment-transaction/core/entity"
)

func (p *PostgresDatabase) findTransaction(ctx context.Context, id string) (*entity.Transaction, error) {
	log.Println("repository.transactions.findTransaction")
	query := `
	SELECT id, status, value, approved_at, denied_at, created_at, updated_at, deleted_at
	FROM transactions
	WHERE id = $1`
	stmt, err := p.DB.PrepareContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	rows, err := stmt.QueryContext(ctx, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var tr *dto.TransactionDto
	log.Println("vai executar rows.Next()")
	for rows.Next() {
		log.Println("vai executar rows.Scan")
		err = rows.Scan(&tr.ID, &tr.Status, &tr.Value, &tr.ApprovedAt, &tr.DeniedAt, &tr.CreatedAt, &tr.UpdatedAt, &tr.DeletedAt)
		if err != nil {
			log.Printf("err :%v\n", err)
			return nil, err
		}
	}
	log.Println("Passou pelo rows.Next()")
	if err = rows.Err(); err != nil {
		log.Printf("err :%v\n", err)
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
