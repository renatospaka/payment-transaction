package entity

import (
	"time"

	pkgEntity "github.com/renatospaka/payment-transaction/utils/entity"
)

type TransactionMount struct {
	ID              string
	ClientID        string
	AuthorizationID string
	Value           float32
	Status          string
	DeniedAt        time.Time
	ApprovedAt      time.Time
	*pkgEntity.TrailDate
}

// Recreate an existing transaction
func MountTransaction(mounting *TransactionMount) (*Transaction, error) {
	uuid, err := pkgEntity.Parse(mounting.ID)
	if err != nil {
		return nil, err
	}

	uuidClient, err := pkgEntity.Parse(string(mounting.ClientID))
	if err != nil {
		return nil, err
	}

	uuidCAuthorization, err := pkgEntity.Parse(mounting.AuthorizationID)
	if err != nil {
		return nil, err
	}

	status := mounting.Status
	if status!= TR_APPROVED &&
		status!= TR_DELETED &&
		status!= TR_DENIED &&
		status!= TR_PENDING {
		status = TR_PENDING
	}

	transaction := &Transaction{
		id:              uuid,
		clientId:        uuidClient,
		authorizationId: uuidCAuthorization,
		value:           mounting.Value,
		status:          status,
		deniedAt:        mounting.DeniedAt,
		approvedAt:      mounting.ApprovedAt,
		TrailDate:       &pkgEntity.TrailDate{},
		valid:           false,
	}
	transaction.TrailDate.SetCreationToDate(mounting.CreatedAt())
	transaction.TrailDate.SetAlterationToToday()
	if !mounting.DeletedAt().IsZero() {
		transaction.TrailDate.SetDeletionToDate(mounting.DeletedAt())
	}

	// deliver the new transaction validated
	err = transaction.Validate()
	if err != nil {
		return nil, err
	}
	return transaction, nil
}
