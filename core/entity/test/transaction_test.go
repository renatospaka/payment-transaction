package test

import (
	"testing"
	"time"

	"github.com/renatospaka/payment-transaction/core/entity"
	"github.com/stretchr/testify/assert"
)

func TestNewTransaction(t *testing.T) {
	tr, err := entity.NewTransaction(200.00)
	assert.Nil(t, err)
	assert.NotNil(t, tr)
	assert.NotEmpty(t, tr.GetID())
	assert.Equal(t, float32(200.00), tr.GetValue())
	assert.True(t, tr.DeniedAt().IsZero())
	assert.True(t, tr.ApprovedAt().IsZero())
	assert.Equal(t, entity.TR_PENDING, tr.GetStatus())

	err = tr.Validate()
	assert.Nil(t, err)
}

func TestNewTransactionValueIsZero(t *testing.T) {
	tr, err := entity.NewTransaction(0)
	assert.NotNil(t, err)
	assert.Nil(t, tr)
	assert.EqualError(t, err, entity.ErrValueIsZero.Error())
}

func TestNewTransactionValueIsNegative(t *testing.T) {
	tr, err := entity.NewTransaction(-200.00)
	assert.NotNil(t, err)
	assert.Nil(t, tr)
	assert.EqualError(t, err, entity.ErrValueIsNegative.Error())
}

func TestSetValue(t *testing.T) {
	tr, err := entity.NewTransaction(200.00)
	assert.Nil(t, err)
	assert.NotNil(t, tr)

	err = tr.SetValue(100)
	assert.Nil(t, err)
	assert.Equal(t, float32(100), tr.GetValue())

	err = tr.SetValue(0)
	assert.NotNil(t, err)
	assert.Equal(t, float32(100), tr.GetValue())

	err = tr.SetValue(-200)
	assert.NotNil(t, err)
	assert.Equal(t, float32(100), tr.GetValue())

	err = tr.Validate()
	assert.Nil(t, err)
}

func TestSetStatusToApproved(t *testing.T) {
	tr, err := entity.NewTransaction(200.00)
	assert.Nil(t, err)
	assert.NotNil(t, tr)
	assert.Equal(t, entity.TR_PENDING, tr.GetStatus())
	assert.True(t, tr.DeniedAt().IsZero())
	assert.True(t, tr.ApprovedAt().IsZero())

	tr.SetStatusToApproved()
	assert.Equal(t, entity.TR_APPROVED, tr.GetStatus())
	assert.True(t, tr.DeniedAt().IsZero())
	assert.False(t, tr.ApprovedAt().IsZero())
}

func TestSetStatusToApprovedOnDate(t *testing.T) {
	tr, err := entity.NewTransaction(200.00)
	assert.Nil(t, err)
	assert.NotNil(t, tr)
	assert.Equal(t, entity.TR_PENDING, tr.GetStatus())
	assert.True(t, tr.DeniedAt().IsZero())
	assert.True(t, tr.ApprovedAt().IsZero())

	date := time.Now().Add(time.Hour * 24 * -10)
	tr.SetStatusToApprovedOnDate(date)
	assert.Equal(t, entity.TR_APPROVED, tr.GetStatus())
	assert.True(t, tr.DeniedAt().IsZero())
	assert.False(t, tr.ApprovedAt().IsZero())
	assert.Equal(t, date.Unix(), tr.ApprovedAt().Unix())
}

func TestSetStatusToDenied(t *testing.T) {
	tr, err := entity.NewTransaction(200.00)
	assert.Nil(t, err)
	assert.NotNil(t, tr)
	assert.Equal(t, entity.TR_PENDING, tr.GetStatus())
	assert.True(t, tr.DeniedAt().IsZero())
	assert.True(t, tr.ApprovedAt().IsZero())

	tr.SetStatusToDenied()
	assert.Equal(t, entity.TR_DENIED, tr.GetStatus())
	assert.False(t, tr.DeniedAt().IsZero())
	assert.True(t, tr.ApprovedAt().IsZero())
}

func TestSetStatusToDeniedOnDate(t *testing.T) {
	tr, err := entity.NewTransaction(200.00)
	assert.Nil(t, err)
	assert.NotNil(t, tr)
	assert.Equal(t, entity.TR_PENDING, tr.GetStatus())
	assert.True(t, tr.DeniedAt().IsZero())
	assert.True(t, tr.ApprovedAt().IsZero())

	date := time.Now().Add(time.Hour * 24 * -10)
	tr.SetStatusToDeniedOnDate(date)
	assert.Equal(t, entity.TR_DENIED, tr.GetStatus())
	assert.False(t, tr.DeniedAt().IsZero())
	assert.True(t, tr.ApprovedAt().IsZero())
	assert.Equal(t, date.Unix(), tr.DeniedAt().Unix())
}

func TestSetStatusToDeleted(t *testing.T) {
	tr, err := entity.NewTransaction(200.00)
	assert.Nil(t, err)
	assert.NotNil(t, tr)
	assert.Equal(t, entity.TR_PENDING, tr.GetStatus())
	assert.True(t, tr.DeniedAt().IsZero())
	assert.True(t, tr.ApprovedAt().IsZero())

	tr.SetStatusToDeleted()
	assert.Equal(t, entity.TR_DELETED, tr.GetStatus())
	assert.True(t, tr.DeniedAt().IsZero())
	assert.True(t, tr.ApprovedAt().IsZero())
	assert.False(t, tr.TrailDate.DeletedAt().IsZero())
}
