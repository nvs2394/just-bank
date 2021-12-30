package domain

import (
	"strconv"

	"github.com/jmoiron/sqlx"
	"github.com/nvs2394/just-bank/errs"
	"github.com/nvs2394/just-bank/logger"
)

type TransactionRepositoryDb struct {
	client *sqlx.DB
}

func (transactionRepoDB TransactionRepositoryDb) Save(transaction Transaction) (*Transaction, *errs.AppError) {
	newTransactionSql := "INSERT INTO transactions (transaction_id, account_id, transaction_type, transaction_date, amount) values (?,?,?,?,?)"

	result, err := transactionRepoDB.client.Exec(newTransactionSql, transaction.Id, transaction.AccountId, transaction.TransactionType, transaction.TransactionDate, transaction.Amount)
	if err != nil {
		logger.Error("Error while creating new transaction " + err.Error())
		return nil, errs.NewUnexpectedError("Unexpected database error")
	}

	id, err := result.LastInsertId()
	if err != nil {
		return nil, errs.NewUnexpectedError("Unexpected error from database")
	}

	transaction.Id = strconv.FormatInt(id, 10)
	return &transaction, nil
}

func NewTransactionRepositoryDb(dbClient *sqlx.DB) TransactionRepositoryDb {
	return TransactionRepositoryDb{dbClient}
}
