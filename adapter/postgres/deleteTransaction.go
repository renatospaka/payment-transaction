package repository

import "github.com/renatospaka/payment-transaction/core/dto"

func (p *PostgresDatabase) deleteTransaction(transactionId *dto.TransactionDeleteDto) error {
	return nil
}
