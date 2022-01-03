package domain

import (
	"github.com/nvs2394/just-bank/dto"
	"github.com/nvs2394/just-bank/errs"
)

type Account struct {
	Id          string `db:"account_id"`
	CustomerId  string `db:"customer_id"`
	OpeningDate string `db:"opening_date"`
	AccountType string `db:"account_type"`
	Amount      float64
	Status      string
}

func (account Account) ToNewAccountResponseDto() dto.NewAccountResponse {
	return dto.NewAccountResponse{AccountId: account.Id}
}

func (account Account) CanWithdrawal(amount float64) bool {
	return account.Amount >= amount
}

type AccountRepository interface {
	Save(Account) (*Account, *errs.AppError)
	FindById(id string) (*Account, *errs.AppError)
	SaveTransaction(transaction Transaction) (*Transaction, *errs.AppError)
}
