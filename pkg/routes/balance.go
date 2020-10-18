package routes

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"practice/api/balance"
	"practice/pkg/model"
)

type balanceHandler struct {
	account *model.Account
}

func AddBalanceHandler(e *gin.Engine, account *model.Account) {
	handler := &balanceHandler{
		account: account,
	}

	e.GET("/", handler.Current)
}

func (b *balanceHandler) Current(ctx *gin.Context) {
	response := balance.Balance{
		CurrentAccountBalance: b.account.Balance(),
	}

	ctx.JSON(http.StatusOK, response)
}
