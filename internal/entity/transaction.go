package entity

import (
	"time"

	pkgEntity "github.com/renatospaka/transaction/pkg/entity"
)

const (
	TR_APPROVED = "approved"
	TR_DENIED   = "denied"
	TR_PENDING  = "pending"
	TR_DELETED  = "deleted"
)

type Transaction struct {
	ID          pkgEntity.ID
	Value       float32
	Status      string
	DeniedAt    time.Time
	ApprovedAt  time.Time
	*pkgEntity.TrailDate
}


// Create a new transaction
func NewTransaction(value float32) (transaction *Transaction, err error) {
	transaction = &Transaction{
		ID:          pkgEntity.NewID(),
		Value:       value,
		Status:      TR_PENDING,
		DeniedAt:    time.Time{},
		ApprovedAt:  time.Time{},
	}
	transaction.TrailDate.SetCreationToToday()

	// deliver the new transaction validated
	err = transaction.Validate()
	if err != nil {
		return nil, err
	}
	return transaction, nil
}

// // Mount an existing transaction
// func MountTransaction(transaction )


// Change the status to approved and the approved day is today
func (t *Transaction) SetStatusToApproved() {
	t.ApprovedAt = time.Now()
	t.DeniedAt = time.Time{}
	t.TrailDate.SetAlterationdToToday()
	t.Status = TR_APPROVED
}


// Change the status to denied and the denied day is today
func (t *Transaction) SetStatusToDenied() {
	t.ApprovedAt = time.Time{}
	t.DeniedAt =  time.Now()
	t.TrailDate.SetAlterationdToToday()
	t.Status = TR_DENIED
}


// Change the status to deleted and the deleted day is today
func (t *Transaction) SetStatusToDeleted() {
	t.ApprovedAt = time.Time{}
	t.DeniedAt = time.Time{}
	t.TrailDate.SetDeletionToToday()
	t.Status = TR_DELETED
}


// Change the transaction value and validate it before committing it
func (t *Transaction) ChangeValue(value float32) (err error) {
	current := t.Value
	t.Value = value

	if err = t.Validate(); err != nil {
		t.Value = current
		return err
	}
	return nil
}


// Validates all business rules for this transaction
func (t *Transaction) Validate() (err error) {
	if t.ID.String() == "" {
		return ErrIDIsRequired
	}

	if _, err = pkgEntity.Parse(t.ID.String()); err != nil {
		return ErrInvalidID
	}

	if t.Value < 0 {
		return ErrValueIsNegative
	}

	if t.Value == 0 {
		return ErrValueIsZero
	}	
	
	if t.Status != TR_APPROVED &&
	t.Status != TR_DELETED &&
	t.Status != TR_DENIED &&
	t.Status != TR_PENDING {
		return ErrInvalidStatus
	}

	return nil
}
