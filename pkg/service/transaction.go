package service

import (
	"github.com/google/uuid"
	"practice/api/transaction"
	"practice/pkg/errors"
	"practice/pkg/log"
	"practice/pkg/model"
	"practice/pkg/storage"
	"time"
)

type TransactionService interface {
	Create(create *transaction.TransactionCreate) (*model.Transaction, errors.Error)
	Get(transactionId string) *model.Transaction
	GetAll() []*model.Transaction
}

func NewTransactionService(transactionStorage storage.Transaction, account *model.Account, log log.Logger) TransactionService {
	return &transactionService{
		transactionStorage: transactionStorage,
		log:                log,
		account:            account,
	}
}

type transactionService struct {
	transactionStorage storage.Transaction
	log                log.Logger
	account            *model.Account
}

func (ts *transactionService) Create(transactionCreate *transaction.TransactionCreate) (*model.Transaction, errors.Error) {
	ts.log.Infof("Creating transaction: %v", transactionCreate)

	var newUuid uuid.UUID
	var err error

	if newUuid, err = uuid.NewRandom(); err != nil {
		return nil, errors.New(err.Error())
	}

	transactionModel := &model.Transaction{
		ID:              newUuid.String(),
		TransactionType: transactionCreate.Type,
		Amount:          transactionCreate.Amount,
		CreatedAt:       time.Now(),
	}

	balance := ts.account.LockBalance()
	defer ts.account.UnlockBalance()

	if transactionModel.TransactionType == "credit" {
		balance += transactionModel.Amount
	} else if transactionModel.TransactionType == "debit" {
		balance -= transactionModel.Amount
	}

	ts.account.UpdateBalance(balance)

	ts.transactionStorage.Create(transactionModel)

	return transactionModel, nil
}

func (ts *transactionService) Get(transactionId string) *model.Transaction {
	return ts.transactionStorage.Get(transactionId)
}

func (ts *transactionService) GetAll() []*model.Transaction {
	return ts.transactionStorage.GetAll()
}
