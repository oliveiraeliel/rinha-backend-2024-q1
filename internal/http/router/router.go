package router

import (
	"github.com/gin-gonic/gin"
	"github.com/oliveiraeliel/rinha-backend-24-q1/internal/http/controller"
	"github.com/oliveiraeliel/rinha-backend-24-q1/internal/transaction"
)

func Run(transactionRepository transaction.TransactionRepository) {
	router := gin.Default()

	transactionService := transaction.NewTransactionService(transactionRepository)
	transactionController := controller.NewTransactionController(transactionService)

	v := router.Group("/api/cliente/:id")
	{
		v.POST("/transacao", transactionController.CreateTransactionHandler)
		v.GET("/extrato", transactionController.GetExtractHandler)
	}

	router.Run()
}
