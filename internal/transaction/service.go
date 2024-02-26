package transaction

import (
	"context"

	"github.com/oliveiraeliel/rinha-backend-24-q1/internal/domain"
)

type TransactionService interface {
	CreateTransaction(ctx context.Context, t domain.Transaction) (domain.CreateTransactionResponse, error)
	GenerateExtract(ctx context.Context, clientId int) (domain.Extract, error)
}

type transactionService struct {
	repository TransactionRepository
}

func NewTransactionService(r TransactionRepository) TransactionService {
	return &transactionService{
		repository: r,
	}
}

func (s *transactionService) CreateTransaction(ctx context.Context, t domain.Transaction) (domain.CreateTransactionResponse, error) {
	return s.repository.SaveTransaction(ctx, &t)
}

func (s *transactionService) GenerateExtract(ctx context.Context, clientId int) (domain.Extract, error) {
	return s.repository.GetExtract(ctx, clientId)
}
