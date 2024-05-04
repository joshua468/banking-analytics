package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/joshua468/banking-analytics/service"
)

type TransactionAPI struct {
	transactionService service.TransactionService
}

func NewTransactionAPI(transactionService service.TransactionService) *TransactionAPI {
	return &TransactionAPI{transactionService: transactionService}
}

func (api *TransactionAPI) RegisterRoutes(router *gin.RouterGroup) {
	router.GET("/transactions", api.GetTransactions)
	// Add more API endpoints here as needed
}

func (api *TransactionAPI) GetTransactions(c *gin.Context) {
	// Implement logic to fetch transactions from service
	transactions, err := api.transactionService.GetTransactions(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, transactions)
}
