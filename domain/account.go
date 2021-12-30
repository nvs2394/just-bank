package domain

import (
	"github.com/nvs2394/just-bank/dto"
	"github.com/nvs2394/just-bank/errs"
)

type Account struct {
	Id          string `db:"account_id"`
	CustomerId  string `db:"customer_id"`
	OpeningDate string `db:"opening_date"`
	AccountType string `db:"account_Type"`
	Amount      float64
	Status      string
}

func (account Account) ToNewAccountResponseDto() dto.NewAccountResponse {
	return dto.NewAccountResponse{AccountId: account.Id}
}

type AccountRepository interface {
	Save(Account) (*Account, *errs.AppError)
}
