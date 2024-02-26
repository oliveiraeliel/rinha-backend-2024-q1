package controller

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/oliveiraeliel/rinha-backend-24-q1/internal/domain"
	"github.com/oliveiraeliel/rinha-backend-24-q1/internal/transaction"
)

type TransactionController struct {
	service transaction.TransactionService
}

type TransactionBodyRequest struct {
	Value       int64  `json:"valor" validate:"required,gt=0"`
	Type        string `json:"tipo" validate:"required,len=1"`
	Description string `json:"descricao" validate:"required,len=10"`
}

func (t *TransactionBodyRequest) Validate() error {
	return nil
}

func NewTransactionController(s transaction.TransactionService) TransactionController {
	return TransactionController{service: s}
}

func (t *TransactionController) CreateTransactionHandler(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "Parâmetro não é um inteiro",
		})
		return
	}

	request := TransactionBodyRequest{}
	err = ctx.ShouldBindJSON(&request)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Body request inválido",
			"error":   err,
		})
		return
	}

	transaction := domain.Transaction{
		ClientID:    id,
		Value:       int(request.Value),
		Type:        request.Type,
		Description: request.Description,
	}

	response, err := t.service.CreateTransaction(ctx, transaction)

	if err != nil {
		if err.Error() == "saldo insuficiente" {
			ctx.JSON(http.StatusUnprocessableEntity, gin.H{
				"message": "Saldo insuficiente",
			})
		} else {
			ctx.JSON(http.StatusNotFound, gin.H{
				"message": "Cliente não cadastrado",
			})
		}
		return
	}

	ctx.JSON(http.StatusOK, response)
}

func (t *TransactionController) GetExtractHandler(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "Parâmetro não é um inteiro",
		})
		return
	}

	extract, err := t.service.GenerateExtract(ctx, id)

	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"message": "Cliente não cadastrado",
		})
		return
	}

	ctx.JSON(http.StatusOK, extract)
}
