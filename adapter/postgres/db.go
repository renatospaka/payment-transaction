package repository

import (
	"context"
	"database/sql"
	"time"

	_ "github.com/lib/pq"

	"github.com/renatospaka/payment-transaction/core/entity"
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

// func connect() *sql.DB {
// 	configs, err := configs.LoadConfig("../../app")
// 	if err != nil {
// 		log.Panic(err)
// 	}

// 	//open connection to the database
// 	conn := "postgresql://" + configs.DBUser + ":" + configs.DBPassword + "@" + configs.DBHost + "/" + configs.DBName + "?sslmode=disable"
// 	db, err := sql.Open("postgres", conn)
// 	if err != nil {
// 		log.Panic(err)
// 	}
// 	defer db.Close()

// 	return db
// }

// Add context to the methodo (in near future)
func (p *PostgresDatabase) Create(transaction *entity.Transaction) error {
	ctx := context.Background()
	return p.createTransaction(ctx, transaction)
}

// Add context to the methodo (in near future)
func (p *PostgresDatabase) Delete(transactionId string) error {
	panic("implement me")
	// return p.deleteTransaction(transactionId)
}

// Add context to the methodo (in near future)
func (p *PostgresDatabase) Update(transaction *entity.Transaction) error {
	panic("implement me")
	// return p.updateTransaction(transaction)
}

// Add context to the methodo (in near future)
func (p *PostgresDatabase) Find(transactionId string) (*entity.Transaction, error) {
	panic("implement me")
	// 	return p.findTransaction(transactionId)
}

func (p *PostgresDatabase) FindAll(page, limit int, sort string) ([]*entity.Transaction, error) {
	panic("implement me")
	// return p.findAllTransactions(page, limit, sort)
}


func isNullDate(date time.Time) bool {
	if date.IsZero() {
		return true
	}
	return false
}