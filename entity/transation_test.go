package entity_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/vitorath/clean-architecture-golang/entity"
)

func TestTransactionWithAmountGreaterThan1000(t *testing.T) {
	transaction := entity.NewTransaction()
	transaction.ID = "1"
	transaction.AccountID = "1"
	transaction.Amount = 2000
	err := transaction.IsValid()
	assert.Error(t, err)
	assert.Equal(t, "you don't have limit for this transaction", err.Error())
}
