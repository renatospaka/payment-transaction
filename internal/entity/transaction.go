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
	deniedAt   time.Time
	approvedAt time.Time
	*pkgEntity.TrailDate
	id         pkgEntity.ID
	status     string
	Value      float32
}

// type Authorization struct {
// 	id          pkgEntity.ID
// 	processedAt time.Time
// 	status      string
// 	value       float32
// }

// Create a new transaction
func NewTransaction(value float32) (*Transaction, error) {
	transaction := &Transaction{
		id:         pkgEntity.NewID(),
		Value:      value,
		status:     TR_PENDING,
		deniedAt:   time.Time{},
		approvedAt: time.Time{},
		TrailDate:  &pkgEntity.TrailDate{},
	}
	transaction.TrailDate.SetCreationToToday()

	// deliver the new transaction validated
	err := transaction.Validate()
	if err != nil {
		return nil, err
	}
	return transaction, nil
}

// Get the ID of the transaction
func (t *Transaction) GetID() string {
	return t.id.String()
}

// Change the status to approved and the approved day is today
func (t *Transaction) SetStatusToApproved() {
	t.approvedAt = time.Now()
	t.deniedAt = time.Time{}
	t.TrailDate.SetAlterationToToday()
	t.status = TR_APPROVED
}

// Change the status to approved on a specific date
func (t *Transaction) SetStatusToApprovedOnDate(date time.Time) {
	t.approvedAt = date
	t.deniedAt = time.Time{}
	t.TrailDate.SetAlterationToToday()
	t.status = TR_APPROVED
}

// Change the status to denied and the denied day is today
func (t *Transaction) SetStatusToDenied() {
	t.approvedAt = time.Time{}
	t.deniedAt = time.Now()
	t.TrailDate.SetAlterationToToday()
	t.status = TR_DENIED
}

// Change the status to denied on a specific date
func (t *Transaction) SetStatusToDeniedOnDate(date time.Time) {
	t.approvedAt = time.Time{}
	t.deniedAt = date
	t.TrailDate.SetAlterationToToday()
	t.status = TR_DENIED
}

// Change the status to deleted and the deleted day is today
func (t *Transaction) SetStatusToDeleted() {
	t.approvedAt = time.Time{}
	t.deniedAt = time.Time{}
	t.TrailDate.SetDeletionToToday()
	t.status = TR_DELETED
}

// Get the current status of the transacion
func (t *Transaction) GetStatus() string {
	return t.status
}

// Get when the transacion was denied (if it was)
func (t *Transaction) DeniedAt() time.Time {
	return t.deniedAt
}

// Get when the transacion was approved (if it was)
func (t *Transaction) ApprovedAt() time.Time {
	return t.approvedAt
}

// Change the transaction value and validate it before committing it
func (t *Transaction) SetValue(value float32) (err error) {
	current := t.Value
	t.Value = value

	if err = t.Validate(); err != nil {
		t.Value = current
		return err
	}
	return nil
}

// Get the current value of the transaction
func (t *Transaction) GetValue() float32 {
	return t.Value
}

// Validates all business rules for this transaction
func (t *Transaction) Validate() error {
	if t.id.String() == "" {
		return ErrIDIsRequired
	}

	if _, err := pkgEntity.Parse(t.id.String()); err != nil {
		return ErrInvalidID
	}

	if t.Value < 0 {
		return ErrValueIsNegative
	}

	if t.Value == 0 {
		return ErrValueIsZero
	}

	if t.status != TR_APPROVED &&
		t.status != TR_DELETED &&
		t.status != TR_DENIED &&
		t.status != TR_PENDING {
		return ErrInvalidStatus
	}

	return nil
}
