package repository

import "github.com/renatospaka/payment-transaction/core/dto"

func (p *PostgresDatabase) findAllTransactions(page, limit int, sort string) ([]*dto.TransactionDto, error) {
	return nil, nil
}
