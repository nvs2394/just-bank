package dto

import (
	"strings"

	"github.com/nvs2394/just-bank/errs"
)

type NewTransactionRequest struct {
	AccountId       int     `json:"account_id"`
	TransactionType string  `json:"transaction_type"`
	Amount          float64 `json:"amount"`
}

func (request NewTransactionRequest) Validate() *errs.AppError {
	if request.Amount < 1 {
		return errs.NewBadRequestError("Amount can not less than 1 dollar")
	}

	if strings.ToLower(request.TransactionType) != "withdrawal" && strings.ToLower(request.TransactionType) != "deposit" {
		return errs.NewBadRequestError("Transaction type should be withdrawal or deposit")
	}

	return nil
}
