package service

import (
	"github.com/stretchr/testify/assert"
	"practice/api/transaction"
	"practice/pkg/log"
	"practice/pkg/model"
	"practice/pkg/storage"
	"testing"
)

var creditTransactionAPI = &transaction.TransactionCreate{
	Amount: 100,
	Type:   "credit",
}

var creditTransaction = &model.Transaction{
	TransactionType: "credit",
	Amount:          100,
}

func TestTransactionService(t *testing.T) {
	transactionStorage := storage.NewTransaction()

	account := model.NewAccount()

	service := NewTransactionService(transactionStorage, account, log.NewConfigless())

	response, err := service.Create(creditTransactionAPI)
	creditTransaction.ID = response.ID
	creditTransaction.CreatedAt = response.CreatedAt

	assert.NoError(t, err)
	assert.Equal(t, response, creditTransaction)

	response = service.Get(creditTransaction.ID)
	assert.Equal(t, response, creditTransaction)

	all := service.GetAll()
	assert.Equal(t, all, []*model.Transaction{creditTransaction})
}
