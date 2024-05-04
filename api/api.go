package api

import (
	"github.com/gin-gonic/gin"
	"github.com/joshua468/banking-analytics/service"
)

func InitializeAPI(router *gin.Engine, transactionService service.TransactionService) {
	api := NewTransactionAPI(transactionService)
	api.RegisterRoutes(router.Group("/api/v1"))
}
