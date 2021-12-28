package app

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/nvs2394/just-bank/service"
)

type Customer struct {
	Name    string `json:"full_name"`
	City    string `json:"city"`
	Zipcode string `json:"zip_code"`
}

type CustomerHandlers struct {
	service service.CustomerService
}

func (customerHandler *CustomerHandlers) getCustomers(rw http.ResponseWriter, r *http.Request) {
	customers, _ := customerHandler.service.GetAllCustomer()
	rw.Header().Add("Content-type", "application/json")

	json.NewEncoder(rw).Encode(customers)
}

func (customerHandler *CustomerHandlers) getCustomer(rw http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	customerId := vars["customer_id"]

	customer, err := customerHandler.service.GetCustomer(customerId)

	if err != nil {
		rw.WriteHeader(http.StatusNotFound)
		json.NewEncoder(rw).Encode(err)
	} else {
		rw.Header().Add("Content-type", "application/json")
		json.NewEncoder(rw).Encode(customer)
	}

}

func createCustomer(rw http.ResponseWriter, r *http.Request) {

	fmt.Println("Calling to create customer")
}
