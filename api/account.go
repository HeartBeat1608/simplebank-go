package api

import (
	"net/http"

	db "github.com/Heartbeat1608/simplebank/db/sqlc"
	"github.com/gin-gonic/gin"
)

type CreateAccountRequest struct {
	Owner    string `json:"owner" binding:"required"`
	Currency string `json:"currency" binding:"required,oneof=USD EUR INR"`
}

func (server *Server) createAccount(ctx *gin.Context) {
	var req CreateAccountRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.CreateAccountParams{
		Owner:    req.Owner,
		Currency: req.Currency,
		Balance:  0,
	}

	acc, err := server.store.CreateAccount(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse((err)))
		return
	}

	ctx.JSON(http.StatusCreated, acc)
}
