package routes

import (
	"github.com/gin-gonic/gin"
	"practice/pkg/model"
	"practice/pkg/service"
)

type HttpRoutes struct {
	transactionService service.TransactionService
	account            *model.Account
}

func NewHttpRoutes(transactionService service.TransactionService, account *model.Account) *HttpRoutes {
	return &HttpRoutes{
		transactionService: transactionService,
		account:            account,
	}
}

func (r *HttpRoutes) AddHttpRoutes(e *gin.Engine) {
	addHealthCheckRoutes(e)
	r.addApplicationRoutes(e)
}

func (r *HttpRoutes) addApplicationRoutes(e *gin.Engine) {
	AddTransactionHandler(e, r.transactionService)
	AddBalanceHandler(e, r.account)
}
