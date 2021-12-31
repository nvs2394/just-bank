package service

import (
	"time"

	"github.com/nvs2394/just-bank/domain"
	"github.com/nvs2394/just-bank/dto"
	"github.com/nvs2394/just-bank/errs"
)

type TransactionService interface {
	NewTransaction(dto.NewTransactionRequest) (*dto.NewTransactionResponse, *errs.AppError)
}
type DefaultTransactionService struct {
	repo domain.TransactionRepository
}

func (transactionService DefaultTransactionService) NewTransaction(req dto.NewTransactionRequest) (*dto.NewTransactionResponse, *errs.AppError) {
	if err := req.Validate(); err != nil {
		return nil, err
	}

	data := domain.Transaction{
		Id:              "",
		AccountId:       req.AccountId,
		TransactionType: req.TransactionType,
		Amount:          req.Amount,
		TransactionDate: time.Now().Format("2006-01-02 15:04:05"),
	}

	newTransaction, err := transactionService.repo.Save(data)
	if err != nil {
		return nil, err
	}

	response := newTransaction.ToTransactionResponseDto()
	return &response, nil
}

func NewTransactionService(repo domain.TransactionRepository) DefaultTransactionService {
	return DefaultTransactionService{repo}
}
