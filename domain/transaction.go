package domain

import (
	"github.com/nvs2394/just-bank/dto"
	"github.com/nvs2394/just-bank/errs"
)

type Transaction struct {
	Id              string `db:"transaction_id"`
	AccountId       string `db:"account_id"`
	TransactionType string `db:"transaction_type"`
	TransactionDate string `db:"transaction_date"`
	Amount          float64
}

func (account Account) ToTransactionResponseDto() dto.NewAccountResponse {
	return dto.NewAccountResponse{AccountId: account.Id}
}

type TransactionRepository interface {
	Save(Transaction) (*Transaction, *errs.AppError)
}
