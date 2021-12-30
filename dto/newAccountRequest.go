package dto

import (
	"strings"

	"github.com/nvs2394/just-bank/errs"
)

type NewAccountRequest struct {
	CustomerId  string  `json:"customer_id"`
	AccountType string  `json:"account_type"`
	Amount      float64 `json:"amount"`
}

func (request NewAccountRequest) Validate() *errs.AppError {
	if request.Amount < 5000 {
		return errs.NewBadRequestError("To open a new account you need to deposit atleast 5000.00")
	}

	if strings.ToLower(request.AccountType) != "saving" && strings.ToLower(request.AccountType) != "checking" {
		return errs.NewBadRequestError("Account type should be checking or saving")
	}
	return nil
}
