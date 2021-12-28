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

func (customerHandler *CustomerHandlers) getCustomers(rw http.ResponseWriter, r *http.Request) {
	status := r.URL.Query().Get("status")

	customers, err := customerHandler.service.GetAllCustomer(status)

	if err != nil {
		writeResponse(rw, err.Code, err.AsMessage())
	} else {
		writeResponse(rw, http.StatusOK, customers)
	}

}

func (customerHandler *CustomerHandlers) getCustomer(rw http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	customerId := vars["customer_id"]

	customer, err := customerHandler.service.GetCustomer(customerId)

	if err != nil {
		writeResponse(rw, err.Code, err.AsMessage())
	} else {
		writeResponse(rw, http.StatusOK, customer)
	}

}

func createCustomer(rw http.ResponseWriter, r *http.Request) {
	fmt.Println("Calling to create customer")
}

func writeResponse(rw http.ResponseWriter, code int, data interface{}) {
	rw.Header().Add("Content-type", "application/json")
	rw.WriteHeader(code)
	if err := json.NewEncoder(rw).Encode(data); err != nil {
		panic(err)
	}
}
