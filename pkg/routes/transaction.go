package routes

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"practice/api/transaction"
	"practice/pkg/errors"
	"practice/pkg/model"
	"practice/pkg/server/context"
	"practice/pkg/service"
)

type transactionHandler struct {
	transactionService service.TransactionService
}

func AddTransactionHandler(e *gin.Engine, service service.TransactionService) {
	handler := &transactionHandler{
		transactionService: service,
	}

	e.POST("/transactions", handler.Create)
	e.GET("/transactions", handler.GetAll)
	e.GET("/transactions/:id", handler.Get)
}

func (t *transactionHandler) Create(ctx *gin.Context) {
	if err := context.CheckContentType(ctx); err != nil {
		ctx.Error(err)
		return
	}

	var transactionCreate transaction.TransactionCreate
	var transactionId string
	if err := context.DecodeBody(ctx, &transactionCreate); err != nil {
		ctx.Error(err)
		return
	}

	if response, err := t.transactionService.Create(&transactionCreate); err != nil {
		ctx.Error(err)
		return
	} else {
		transactionId = response.ID
	}

	ctx.JSON(http.StatusOK, transactionId)
}

func (t *transactionHandler) Get(ctx *gin.Context) {
	transactionId := ctx.Param("id")

	if modelTransaction := t.transactionService.Get(transactionId); modelTransaction != nil {
		response := t.transactionToApi(modelTransaction)
		ctx.JSON(http.StatusOK, response)
	} else {
		ctx.Error(errors.NotFound.Newf("transaction %s not found", transactionId))
	}
}

func (t *transactionHandler) GetAll(ctx *gin.Context) {
	var response []transaction.Transaction

	for _, transact := range t.transactionService.GetAll() {
		response = append(response, t.transactionToApi(transact))
	}

	ctx.JSON(http.StatusOK, response)
}

func (t *transactionHandler) transactionToApi(model *model.Transaction) transaction.Transaction {
	return transaction.Transaction{
		Id:     model.ID,
		Amount: model.Amount,
		Type:   model.TransactionType,
		Date:   model.CreatedAt,
	}
}
