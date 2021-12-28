package app

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/nvs2394/just-bank/service"
)

type CustomerHandlers struct {
	service service.CustomerService
}

func (customerHandler *CustomerHandlers) getCustomers(response http.ResponseWriter, r *http.Request) {
	status := r.URL.Query().Get("status")

	customers, err := customerHandler.service.GetAllCustomer(status)

	if err != nil {
		writeResponse(response, err.Code, err.AsMessage())
	} else {
		writeResponse(response, http.StatusOK, customers)
	}

}

func (customerHandler *CustomerHandlers) getCustomer(response http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	customerId := vars["customer_id"]

	customer, err := customerHandler.service.GetCustomer(customerId)

	if err != nil {
		writeResponse(response, err.Code, err.AsMessage())
	} else {
		writeResponse(response, http.StatusOK, customer)
	}

}

func createCustomer(response http.ResponseWriter, r *http.Request) {
	fmt.Println("Calling to create customer")
}

func writeResponse(response http.ResponseWriter, code int, data interface{}) {
	response.Header().Add("Content-type", "application/json")
	response.WriteHeader(code)
	if err := json.NewEncoder(response).Encode(data); err != nil {
		panic(err)
	}
}
