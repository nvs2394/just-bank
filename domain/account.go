package domain

import (
	"github.com/nvs2394/just-bank/errs"
)

type Account struct {
	Id          string `db:"account_id"`
	CustomerId  string `db:"customer_id"`
	OpeningDate string `db:"opening_date"`
	AccountType string `db:"account_Type"`
	Amount      string
	Status      string
}

type AccountRepository interface {
	Save(Account) (*Account, *errs.AppError)
}
