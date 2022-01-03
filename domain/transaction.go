package domain

import (
	"github.com/nvs2394/just-bank/dto"
)

const WITHDRAWAL = "withdrawal"

type Transaction struct {
	Id              string `db:"transaction_id"`
	AccountId       string `db:"account_id"`
	TransactionType string `db:"transaction_type"`
	TransactionDate string `db:"transaction_date"`
	Amount          float64
}

func (transaction Transaction) ToTransactionResponseDto() dto.NewTransactionResponse {
	return dto.NewTransactionResponse{TransactionId: transaction.Id}
}

func (transaction Transaction) IsWithdrawal() bool {
	return transaction.TransactionType == WITHDRAWAL
}

func (transaction Transaction) ToDto() dto.NewTransactionResponse {
	return dto.NewTransactionResponse{
		TransactionId:   transaction.Id,
		AccountId:       transaction.AccountId,
		Amount:          transaction.Amount,
		TransactionType: transaction.TransactionType,
		TransactionDate: transaction.TransactionDate,
	}
}
