package domain

import (
	"database/sql"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/nvs2394/just-bank/errs"
	"github.com/nvs2394/just-bank/logger"
)

type CustomerRepositoryDb struct {
	client *sql.DB
}

func (d CustomerRepositoryDb) FindAll(status string) ([]Customer, *errs.AppError) {
	var rows *sql.Rows
	var err error
	if status == "" {
		findAllSql := "select customer_id, name, city, zipcode,status, date_of_birth from customers"

		rows, err = d.client.Query(findAllSql)

	} else {
		findAllByStatusSql := "select customer_id, name, city, zipcode,status, date_of_birth from customers where status=?"

		rows, err = d.client.Query(findAllByStatusSql, status)
	}

	if err != nil {
		logger.Error("Error while querying customer table " + err.Error())
		return nil, errs.NewUnexpectedError("unexpected database error")
	}

	var customers []Customer

	for rows.Next() {
		var customer Customer
		err := rows.Scan(&customer.Id, &customer.Name, &customer.City, &customer.Zipcode, &customer.Status, &customer.DateOfBirth)
		if err != nil {
			logger.Error("Error while scaning customer table " + err.Error())
			return nil, errs.NewUnexpectedError("unexpected database error")
		}

		customers = append(customers, customer)

	}
	return customers, nil
}

func (d CustomerRepositoryDb) FindById(id string) (*Customer, *errs.AppError) {
	customerSql := "select customer_id, name, city, zipcode, date_of_birth, status from customers where customer_id=?"

	row := d.client.QueryRow(customerSql, id)
	var customer Customer
	err := row.Scan(&customer.Id, &customer.Name, &customer.City, &customer.Zipcode, &customer.Status, &customer.DateOfBirth)

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
	//TODO: will be replate by run time env
	client, err := sql.Open("mysql", "just_bank_user:password@/just_bank_db")

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
