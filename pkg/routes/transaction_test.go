package routes

import (
	"github.com/golang/mock/gomock"
	"practice/pkg/routes/mocks"
)

//go:generate mockgen -destination=mocks/mock_TransactionService.go -package=mocks -source=../service/transaction.go

type transactionMocks struct {
	ctrl    *gomock.Controller
	service *mocks.MockTransactionService
}

/*
func (builder *transactionMocks) build() *gin.Engine {
	router := test_fixture.SetupRouter(log.NewConfigless())
	AddtransactionHandler(router, builder.service)
	return router
}

func transactionSetUp(t *testing.T) (*gin.Engine, *transactionMocks) {
	ctrl := gomock.NewController(t)
	mocks := &transactionMocks{
		ctrl:    ctrl,
		service: mocks.NewMocktransaction(ctrl),
	}
	return mocks.build(), mocks
}
*/
