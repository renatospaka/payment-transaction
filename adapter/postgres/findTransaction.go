package repository

import (
	"context"

	"github.com/renatospaka/payment-transaction/core/dto"
)

func (p *PostgresDatabase) findTransaction(transactionId *dto.TransactionFindDto) (*dto.TransactionDto, error) {
	ctx := context.Background()
	query := `
	SELECT id, status, value, approved_at, denied_a, created_at, updated_at, deleted_at
	FROM transactions
	WHERE id = ?`
	// query := `
	// UPDATE transactions
	// SET status = ?,
	// 	value = ?,
	// 	approved_at = ?,
	// 	denied_at = ?,
	// 	created_at = ?,
	// 	updated_at = ?,
	// 	deleted_at = ?
	// WHERE id = ?
	// )`
	stmt, err := p.DB.PrepareContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	rows, err := stmt.QueryContext(ctx, transactionId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var tr *dto.TransactionDto
	for rows.Next() {
		if err = rows.Scan(tr); err != nil {
			return nil, err
		}
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}

	return tr, nil
}
