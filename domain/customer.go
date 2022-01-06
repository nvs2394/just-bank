package domain

import (
	"github.com/nvs2394/just-bank-lib/errs"
	"github.com/nvs2394/just-bank/dto"
)

type Customer struct {
	Id          string `db:"customer_id"`
	Name        string
	City        string
	Zipcode     string
	DateOfBirth string `db:"date_of_birth"`
	Status      string
}

func (customer Customer) StatusAsText() string {

	statusAsText := "active"
	if customer.Status == "0" {
		statusAsText = "inactive"
	}
	return statusAsText
}

func (customer Customer) ToDto() dto.CustomerResponse {

	return dto.CustomerResponse{
		Id:          customer.Id,
		Name:        customer.Name,
		City:        customer.City,
		Zipcode:     customer.Zipcode,
		DateOfBirth: customer.DateOfBirth,
		Status:      customer.StatusAsText(),
	}
}

type CustomerRepository interface {
	FindAll(status string) ([]Customer, *errs.AppError)
	FindById(string) (*Customer, *errs.AppError)
}
