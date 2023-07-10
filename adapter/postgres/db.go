package repository

import (
	"context"
	"database/sql"

	_ "github.com/lib/pq"

	"github.com/renatospaka/payment-transaction/core/entity"
)

type PostgresDatabase struct {
	DB *sql.DB
}

var ctx context.Context

func NewPostgresDatabase(db *sql.DB) *PostgresDatabase {
	ctx = context.Background()
	return &PostgresDatabase{
		DB: db,
	}
}

// Persist the created transaction into the Postgres
func (p *PostgresDatabase) Create(transaction *entity.Transaction) error {
	return p.createTransaction(ctx, transaction)
}

// Persist a reprocessed transaction into the Postgres
func (p *PostgresDatabase) Reprocess(transaction *entity.Transaction) error {
	return p.reprocessTransaction(ctx, transaction)
}

// Find the transaction by its id
func (p *PostgresDatabase) Find(id string) (*entity.Transaction, error) {
	return p.findById(ctx, id)
}

func (p *PostgresDatabase) FindAll(page, limit int) ([]*entity.Transaction, error) {
	return p.findAllTransactions(ctx, page, limit)
}

func (p *PostgresDatabase) Approve(transaction *entity.Transaction) error {
	return p.approveTransaction(ctx, transaction)
}

func (p *PostgresDatabase) Deny(transaction *entity.Transaction) error {
	return p.denyTransaction(ctx, transaction)
}

// Add context to the methodo (in near future)
func (p *PostgresDatabase) Delete(id string) error {
	return p.deleteTransaction(ctx, id)
}

// Add context to the methodo (in near future)
func (p *PostgresDatabase) Update(tr *entity.Transaction) error {
	return p.updateTransaction(ctx, tr)
}
