package test

import (
	"testing"

	"github.com/renatospaka/transaction/internal/entity"
	"github.com/stretchr/testify/assert"
)

func TestNewTransaction(t *testing.T) {
	tr, err := entity.NewTransaction(200.00)
	assert.Nil(t, err)
	assert.NotNil(t, tr)
	assert.NotEmpty(t, tr.GetID())
	assert.Equal(t, float32(200.00), tr.GetValue())
	assert.True(t, tr.Denied().IsZero())
	assert.True(t, tr.Approved().IsZero())
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