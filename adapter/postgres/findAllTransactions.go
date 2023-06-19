package repository

import (
	"context"
	"database/sql"
	"log"

	"github.com/renatospaka/payment-transaction/core/entity"
	"github.com/renatospaka/payment-transaction/utils/dateTime"
	utils "github.com/renatospaka/payment-transaction/utils/entity"
)

func (p *PostgresDatabase) findAllTransactions(ctx context.Context, page int, limit int) ([]*entity.Transaction, error) {
	log.Println("repository.transactions.findAllTransactions")

	query := `
	SELECT id, client_id, authorization_id, status, value , approved_at, denied_at, created_at, updated_at, deleted_at
	FROM transactions
	OFFSET $2 LIMIT $1`
	stmt, err := p.DB.PrepareContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	offset := (page - 1) * limit
	rows, err := stmt.QueryContext(ctx, limit, offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	// log.Printf("repository.transactions.findAllTransactions - limit: %d, offset: %d\n", limit, offset)
	var (
		transactions  []*entity.Transaction
		transaction   *entity.Transaction
		mounting      *entity.TransactionMount
		approved      sql.NullTime
		denied        sql.NullTime
		created       sql.NullTime
		updated       sql.NullTime
		deleted       sql.NullTime
		id            sql.NullString
		client        sql.NullString
		authorization sql.NullString
		status        sql.NullString
		value         float32
	)
	for rows.Next() {
		err = rows.Scan(
			&id,
			&client,
			&authorization,
			&status,
			&value,
			&approved,
			&denied,
			&created,
			&updated,
			&deleted,
		)
		if err != nil {
			if err == sql.ErrNoRows {
				// log.Printf("repository.transactions.findAllTransactions - deu ruim: %d\n", 1)
				return []*entity.Transaction{}, nil
			}
			// log.Printf("repository.transactions.findAllTransactions - deu ruim: %d\n", 2)
			return []*entity.Transaction{}, err
		}

		approvedAt := dateTime.FormatNullDate(approved)
		deniedAt := dateTime.FormatNullDate(denied)
		createdAt := dateTime.FormatNullDate(created)
		updatedAt := dateTime.FormatNullDate(updated)
		deletedAt := dateTime.FormatNullDate(deleted)
		id, _ := utils.Parse(dateTime.FormatNullString(id))
		clientId, _ := utils.Parse(dateTime.FormatNullString(client))
		authorizationId, _ := utils.Parse(dateTime.FormatNullString(authorization))
		mounting = &entity.TransactionMount{
			ID:              id.String(),
			ClientID:        clientId.String(),
			AuthorizationID: authorizationId.String(),
			Value:           value,
			Status:          dateTime.FormatNullString(status),
			DeniedAt:        deniedAt,
			ApprovedAt:      approvedAt,
			TrailDate:       &utils.TrailDate{},
		}
		mounting.SetCreationToDate(createdAt)
		mounting.SetAlterationToDate(updatedAt)
		mounting.SetDeletionToDate(deletedAt)

		transaction, err = entity.MountTransaction(mounting)
		if err != nil {
			// log.Printf("repository.transactions.findAllTransactions - deu ruim: %d\n", 3)
			return []*entity.Transaction{}, err
		}
		transactions = append(transactions, transaction)
		// log.Printf("repository.transactions.findAllTransactions - TRANSACTIONS tem %d registros\n",  len(transactions))
	}
	if err := rows.Err(); err != nil {
		// log.Printf("repository.transactions.findAllTransactions - deu ruim: %d\n", 4)
		return []*entity.Transaction{}, err
	}

	// log.Printf("repository.transactions.findAllTransactions - deu BBOOOOMMMM: TRANSACTIONS tem %d registros no TOTAL\n",  len(transactions))
	return transactions, nil
}
