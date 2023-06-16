package repository

import (
	"context"
	"database/sql"
	"time"

	_ "github.com/lib/pq"

	"github.com/renatospaka/payment-transaction/core/entity"
	"github.com/renatospaka/payment-transaction/utils"
)

type PostgresDatabase struct {
	DB *sql.DB
}

func NewPostgresDatabase(db *sql.DB) *PostgresDatabase {
	// db := connect()
	return &PostgresDatabase{
		DB: db,
	}
}

// Add context to the methodo (in near future)
func (p *PostgresDatabase) Create(transaction *entity.Transaction) error {
	ctx := context.Background()
	return p.createTransaction(ctx, transaction)
}

func (p *PostgresDatabase) Approve(transaction *entity.Transaction) error {
	ctx := context.Background()
	return p.approveTransaction(ctx, transaction)
}

func (p *PostgresDatabase) Deny(transaction *entity.Transaction) error {
	ctx := context.Background()
	return p.denyTransaction(ctx, transaction)
}

// Add context to the methodo (in near future)
func (p *PostgresDatabase) Delete(id string) error {
	ctx := context.Background()
	return p.deleteTransaction(ctx, id)
}

// Add context to the methodo (in near future)
func (p *PostgresDatabase) Update(transaction *entity.Transaction) error {
	panic("implement me")
	// return p.updateTransaction(transaction)
}

// Add context to the methodo (in near future)
func (p *PostgresDatabase) Find(id string) (*entity.Transaction, error) {
	ctx := context.Background()
	return p.findTransaction(ctx, id)
}

func (p *PostgresDatabase) FindAll(page, limit int, sort string) ([]*entity.Transaction, error) {
	panic("implement me")
	// return p.findAllTransactions(page, limit, sort)
}

func formatNullDate(nullDate sql.NullTime) time.Time {
	if nullDate.Valid {
		return nullDate.Time
	} 
	return utils.NULL_DATE
}
