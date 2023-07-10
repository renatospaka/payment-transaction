package entity

import "errors"

var (
	ErrTransactionIDIsRequired = errors.New("transaction id is required")
	ErrInvalidTransactionID    = errors.New("invalid transaction id")
	ErrTransactionIDNotFound   = errors.New("transaction not found")
	ErrClientIDIsRequired      = errors.New("client id is required")
	ErrInvalidClientID         = errors.New("invalid client id")
	ErrInvalidAuthorizationID  = errors.New("invalid authorization id")
	ErrValueIsNegative         = errors.New("value must be positive")
	ErrValueIsZero             = errors.New("value must be greater than zero")
	ErrInvalidStatus           = errors.New("invalid status")
	ErrCannotReprocess         = errors.New("transaction cannot be reprocessed, status not allowed")
)
