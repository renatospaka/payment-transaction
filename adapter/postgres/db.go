package repository

import (
	// "context"
	"database/sql"
	"log"

	"github.com/renatospaka/payment-transaction/core/dto"
	"github.com/renatospaka/payment-transaction/utils/configs"
)

type PostgresDatabase struct {
	DB *sql.DB
}

func NewPostgresDatabase() *PostgresDatabase {
	log.Println("iniciando conexão com o banco de dados")
	db := connect()
	return &PostgresDatabase{
		DB: db,
	}
}

func connect() *sql.DB {
	configs, err := configs.LoadConfig("../../app")
	if err != nil {
		log.Panic(err)
	}

	//open connection to the database
	conn := "postgresql://" + configs.DBUser + ":" + configs.DBPassword + "@" + configs.DBHost + "/" + configs.DBName + "?sslmode=disable"
	db, err := sql.Open("postgres", conn)
	if err != nil {
		log.Panic(err)
	}
	defer db.Close()

	return db
}

// Add context to the methodo (in near future)
func (p *PostgresDatabase) Create(transaction *dto.TransactionCreateDto) error {
	panic("implement me")
	// return p.createTransaction(transaction)
}

// Add context to the methodo (in near future)
func (p *PostgresDatabase) Delete(transactionId *dto.TransactionDeleteDto) error {
	panic("implement me")
	// return p.deleteTransaction(transactionId)
}

// Add context to the methodo (in near future)
func (p *PostgresDatabase) Update(transaction *dto.TransactionUpdateDto) error {
	panic("implement me")
	// return p.updateTransaction(transaction)
}

// Add context to the methodo (in near future)
func (p *PostgresDatabase) Find(transactionId *dto.TransactionFindDto) (*dto.TransactionDto, error) {
	panic("implement me")
	// 	return p.findTransaction(transactionId)
}

func (p *PostgresDatabase) FindAll(page, limit int, sort string) ([]*dto.TransactionDto, error) {
	panic("implement me")
	// return p.findAllTransactions(page, limit, sort)
}
