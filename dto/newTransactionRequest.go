package dto

import (
	"strings"

	"github.com/nvs2394/just-bank/errs"
)

const WITHDRAWAL = "withdrawal"
const DEPOSIT = "deposit"

type NewTransactionRequest struct {
	AccountId       string  `json:"account_id"`
	TransactionType string  `json:"transaction_type"`
	Amount          float64 `json:"amount"`
	CustomerId      string  `json:"customer_id"`
	TransactionDate string  `json:"transaction_date"`
}

func (request NewTransactionRequest) Validate() *errs.AppError {
	if request.Amount < 0 {
		return errs.NewBadRequestError("Amount cannot be less than zero")
	}

	if strings.ToLower(request.TransactionType) != WITHDRAWAL && strings.ToLower(request.TransactionType) != DEPOSIT {
		return errs.NewBadRequestError("Transaction type should be withdrawal or deposit")
	}

	return nil
}

func (request NewTransactionRequest) IsTransactionTypeWithdrawal() bool {
	return request.TransactionType == WITHDRAWAL
}
