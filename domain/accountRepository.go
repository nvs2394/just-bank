package domain

import (
	"strconv"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/nvs2394/just-bank/errs"
	"github.com/nvs2394/just-bank/logger"
)

type AccountRepositoryDb struct {
	client *sqlx.DB
}

func (d AccountRepositoryDb) Save(account Account) (*Account, *errs.AppError) {
	newAccountSql := "INSERT INTO accounts (customer_id, opening_date, account_type, amount) values(?,?,?,?)"

	result, err := d.client.Exec(newAccountSql, account.CustomerId, account.OpeningDate, account.AccountType, account.Amount)

	if err != nil {
		logger.Error("Error while creating new account " + err.Error())
		return nil, errs.NewUnexpectedError("Unexpected database error")
	}

	id, err := result.LastInsertId()

	if err != nil {
		return nil, errs.NewUnexpectedError("Unexpected error from database")
	}
	account.Id = strconv.FormatInt(id, 10)
	return &account, nil
}

func NewAccountRepositoryDb(dbClient *sqlx.DB) AccountRepositoryDb {

	return AccountRepositoryDb{
		dbClient,
	}
}
