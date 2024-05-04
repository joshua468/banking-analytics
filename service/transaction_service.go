// service/transaction_service.go
package service

import (
	"github.com/joshua468/banking-analytics/models"
	"github.com/joshua468/banking-analytics/repository"

	"github.com/gin-gonic/gin"
)

type TransactionService interface {
	GetTransactions(c *gin.Context) ([]models.Transaction, error)
}

type transactionService struct {
	transactionRepo repository.TransactionRepository
}

func NewTransactionService(transactionRepo repository.TransactionRepository) TransactionService {
	return &transactionService{
		transactionRepo: transactionRepo,
	}
}

func (s *transactionService) GetTransactions(c *gin.Context) ([]models.Transaction, error) {
	ctx := c.Request.Context()

	// Fetch transactions from repository
	transactions, err := s.transactionRepo.GetTransactions(ctx)
	if err != nil {
		return nil, err
	}

	return transactions, nil
}
