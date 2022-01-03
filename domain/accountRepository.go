package domain

import (
	"database/sql"
	"strconv"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/nvs2394/just-bank/errs"
	"github.com/nvs2394/just-bank/logger"
)

type AccountRepositoryDb struct {
	client *sqlx.DB
}

func (accountRepoDB AccountRepositoryDb) Save(account Account) (*Account, *errs.AppError) {
	newAccountSql := "INSERT INTO accounts (customer_id, opening_date, account_type, amount) values(?,?,?,?)"

	result, err := accountRepoDB.client.Exec(newAccountSql, account.CustomerId, account.OpeningDate, account.AccountType, account.Amount)

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

func (accountRepoDB AccountRepositoryDb) FindById(id string) (*Account, *errs.AppError) {
	findIdSql := "SELECT * FROM accounts where account_id=?"
	var account Account

	err := accountRepoDB.client.Get(&account, findIdSql, id)
	if err != nil {
		if err == sql.ErrNoRows {
			logger.Error("Account not found" + err.Error())
			return nil, errs.NewNotFoundError("Account not found")
		} else {
			logger.Error("Error while getting account info " + err.Error())
			return nil, errs.NewUnexpectedError("Unexpected database error")
		}
	}
	return &account, nil
}

func (accountRepoDB AccountRepositoryDb) SaveTransaction(transaction Transaction) (*Transaction, *errs.AppError) {
	trx, err := accountRepoDB.client.Begin()
	if err != nil {
		logger.Error("Error while starting a new transaction for bank account transaction")
		return nil, errs.NewUnexpectedError("Unexpected database error")
	}

	newTransactionSql := "INSERT INTO transactions (account_id, transaction_type, transaction_date, amount) values (?,?,?,?)"

	result, err := trx.Exec(newTransactionSql, transaction.AccountId, transaction.TransactionType, transaction.TransactionDate, transaction.Amount)

	if err != nil {
		logger.Error("Error while creating new transaction " + err.Error())
		return nil, errs.NewUnexpectedError("Unexpected database error")
	}

	if transaction.IsWithdrawal() {
		updateAccountBalanceWithdrawalTypeSql := "UPDATE accounts SET amount = amount - ? where account_id = ?"
		_, err = trx.Exec(updateAccountBalanceWithdrawalTypeSql, transaction.Amount, transaction.AccountId)
	} else {
		updateAccountBalanceDepositTypeSql := "UPDATE accounts SET amount = amount + ? where account_id = ?"
		_, err = trx.Exec(updateAccountBalanceDepositTypeSql, transaction.Amount, transaction.AccountId)
	}

	if err != nil {
		trx.Rollback()
		logger.Error("Error while saving transaction" + err.Error())
		return nil, errs.NewUnexpectedError("Unexpected database error")
	}

	err = trx.Commit()

	if err != nil {
		trx.Rollback()
		logger.Error("Error while commiting transaction for bank account" + err.Error())
		return nil, errs.NewUnexpectedError("Unexpected database error")
	}

	id, err := result.LastInsertId()
	if err != nil {
		logger.Error("Error while getting the last transaction id" + err.Error())
		return nil, errs.NewUnexpectedError("Unexpected error from database")
	}

	account, appErr := accountRepoDB.FindById(transaction.AccountId)
	if appErr != nil {
		return nil, appErr
	}

	transaction.Id = strconv.FormatInt(id, 10)
	transaction.Amount = account.Amount
	return &transaction, nil
}

func NewAccountRepositoryDb(dbClient *sqlx.DB) AccountRepositoryDb {
	return AccountRepositoryDb{
		dbClient,
	}
}
