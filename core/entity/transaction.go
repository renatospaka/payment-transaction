package entity

import (
	"time"

	pkgEntity "github.com/renatospaka/payment-transaction/utils/entity"
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
	id     pkgEntity.ID
	status string
	value  float32
	valid  bool
}

// Create a new transaction
func NewTransaction(value float32) (*Transaction, error) {
	transaction := &Transaction{
		id:         pkgEntity.NewID(),
		value:      value,
		status:     TR_PENDING,
		deniedAt:   time.Time{},
		approvedAt: time.Time{},
		TrailDate:  &pkgEntity.TrailDate{},
		valid:      false,
	}
	transaction.TrailDate.SetCreationToToday()

	// deliver the new transaction validated
	err := transaction.Validate()
	if err != nil {
		return nil, err
	}
	return transaction, nil
}

func MountTransaction(id, status string, value float32, deniedAt time.Time, approvedAt time.Time, createdAt time.Time, updatedAt time.Time, deletedAt time.Time) (*Transaction, error) {
	uuid, err := pkgEntity.Parse(id)
	if err != nil {
		return nil, err
	}

	transaction := &Transaction{
		deniedAt:   deniedAt,
		approvedAt: approvedAt,
		TrailDate:  &pkgEntity.TrailDate{},
		id:         uuid,
		status:     status,
		value:      value,
		valid:      false,
	}
	transaction.TrailDate.SetCreationToDate(createdAt)
	transaction.TrailDate.SetAlterationToToday()
	if !deletedAt.IsZero() {
		transaction.TrailDate.SetDeletionToDate(deletedAt)
	}

	// deliver the new transaction validated
	err = transaction.Validate()
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
	current := t.value
	t.value = value

	if err = t.Validate(); err != nil {
		t.value = current
		return err
	}
	return nil
}

// Get the current value of the transaction
func (t *Transaction) GetValue() float32 {
	return t.value
}

// Validates all business rules for this transaction
func (t *Transaction) Validate() error {
	t.valid = false
	if t.id.String() == "" {
		return ErrIDIsRequired
	}

	if _, err := pkgEntity.Parse(t.id.String()); err != nil {
		return ErrInvalidID
	}

	if t.value < 0 {
		return ErrValueIsNegative
	}

	if t.value == 0 {
		return ErrValueIsZero
	}

	if t.status != TR_APPROVED &&
		t.status != TR_DELETED &&
		t.status != TR_DENIED &&
		t.status != TR_PENDING {
		return ErrInvalidStatus
	}

	t.valid = true
	return nil
}

// Return whether the structure is valid or not
func (t *Transaction) IsValid() bool {
	return t.valid
}
