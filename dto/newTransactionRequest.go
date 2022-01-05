package dto

import (
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

	if !request.IsTransactionTypeWithdrawal() && !request.IsTransactionTypeDeposit() {
		return errs.NewBadRequestError("Transaction type should be withdrawal or deposit")
	}

	return nil
}

func (request NewTransactionRequest) IsTransactionTypeWithdrawal() bool {
	return request.TransactionType == WITHDRAWAL
}

func (request NewTransactionRequest) IsTransactionTypeDeposit() bool {
	return request.TransactionType == DEPOSIT
}
