package domain

import (
	"github.com/nvs2394/just-bank/dto"
	"github.com/nvs2394/just-bank/errs"
)

type Transaction struct {
	Id              string `db:"transaction_id"`
	AccountId       int    `db:"account_id"`
	TransactionType string `db:"transaction_type"`
	TransactionDate string `db:"transaction_date"`
	Amount          float64
}

func (transaction Transaction) ToTransactionResponseDto() dto.NewTransactionResponse {
	return dto.NewTransactionResponse{TransactionId: transaction.Id}
}

type TransactionRepository interface {
	Save(Transaction) (*Transaction, *errs.AppError)
}
