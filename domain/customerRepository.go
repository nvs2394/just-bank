package domain

import (
	"database/sql"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

type CustomerRepositoryDb struct {
	client *sql.DB
}

func (d CustomerRepositoryDb) FindAll() ([]Customer, error) {

	findAllSql := "select customer_id, name, city, zipcode,status, date_of_birth from customers"

	rows, err := d.client.Query(findAllSql)

	if err != nil {
		log.Println("Error while querying customer table " + err.Error())
		return nil, err
	}

	var customers []Customer

	for rows.Next() {
		var customer Customer
		err := rows.Scan(&customer.Id, &customer.Name, &customer.City, &customer.Zipcode, &customer.Status, &customer.DateOfBirth)
		if err != nil {
			log.Println("Error while scaning customer table " + err.Error())
			return nil, err
		}

		customers = append(customers, customer)

	}
	return customers, nil
}

func (d CustomerRepositoryDb) FindById(id string) (*Customer, error) {
	customerSql := "select customer_id, name, city, zipcode, date_of_birth, status from customers where customer_id=?"

	row := d.client.QueryRow(customerSql, id)
	var customer Customer
	err := row.Scan(&customer.Id, &customer.Name, &customer.City, &customer.Zipcode, &customer.Status, &customer.DateOfBirth)

	if err != nil {
		log.Println("Error while scaning customer table " + err.Error())
		return nil, err
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
