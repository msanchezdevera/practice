package storage

import (
	"github.com/magiconair/properties/assert"
	"practice/pkg/model"
	"testing"
	"time"
)

var creditTransaction = &model.Transaction{
	ID:              "test-id",
	TransactionType: "credit",
	Amount:          100,
	CreatedAt:       time.Time{},
}

func TestTransactionStorage(t *testing.T) {
	transactionStorage := NewTransaction()
	transactionStorage.Create(creditTransaction)

	response := transactionStorage.Get(creditTransaction.ID)
	assert.Equal(t, response, creditTransaction)

	allTransactions := transactionStorage.GetAll()
	assert.Equal(t, allTransactions, []*model.Transaction{creditTransaction})
}
