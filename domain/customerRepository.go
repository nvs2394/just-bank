package domain

import (
	"database/sql"
	"os"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/nvs2394/just-bank/errs"
	"github.com/nvs2394/just-bank/logger"
)

type CustomerRepositoryDb struct {
	client *sqlx.DB
}

func (d CustomerRepositoryDb) FindAll(status string) ([]Customer, *errs.AppError) {
	customers := make([]Customer, 0)
	var err error

	if status == "" {
		findAllSql := "select customer_id, name, city, zipcode,status, date_of_birth from customers"

		err = d.client.Select(&customers, findAllSql)

	} else {
		findAllByStatusSql := "select customer_id, name, city, zipcode,status, date_of_birth from customers where status=?"

		err = d.client.Select(&customers, findAllByStatusSql, status)
	}

	if err != nil {
		logger.Error("Error while querying customer table " + err.Error())
		return nil, errs.NewUnexpectedError("unexpected database error")
	}
	return customers, nil
}

func (d CustomerRepositoryDb) FindById(id string) (*Customer, *errs.AppError) {
	customerSql := "select customer_id, name, city, zipcode, date_of_birth, status from customers where customer_id=?"

	var customer Customer
	err := d.client.Get(&customer, customerSql, id)

	if err != nil {
		if err == sql.ErrNoRows {
			logger.Error("Customer not found" + err.Error())
			return nil, errs.NewNotFoundError("Customer not found")
		} else {
			logger.Error("Error while scaning customer table " + err.Error())
			return nil, errs.NewUnexpectedError("unexpected database error")
		}
	}

	return &customer, nil
}

func NewCustomerRepositoryDb() CustomerRepositoryDb {
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")
	connectionString := dbUser + ":" + dbPassword + "@/" + dbName

	client, err := sqlx.Open("mysql", connectionString)

	if err != nil {
		panic(err)
	}

	client.SetConnMaxLifetime(time.Minute * 3)
	client.SetMaxOpenConns(10)
	client.SetMaxIdleConns(10)
	return CustomerRepositoryDb{
		client,
	}
}
