package app

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

type Customer struct {
	Name    string `json:"full_name"`
	City    string `json:"city"`
	Zipcode string `json:"zip_code"`
}

func getCustomers(rw http.ResponseWriter, r *http.Request) {
	customers := []Customer{
		{Name: "Bao", City: "Hanoi", Zipcode: "110075"},
		{Name: "Sonny", City: "Hanoi", Zipcode: "110076"},
	}
	rw.Header().Add("Content-type", "application/json")

	json.NewEncoder(rw).Encode(customers)
}

func getCustomer(rw http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	fmt.Println(vars["customer_id"])
	customer := Customer{
		Name: "Bao", City: "Hanoi", Zipcode: "110075",
	}
	rw.Header().Add("Content-type", "application/json")

	json.NewEncoder(rw).Encode(customer)
}

func createCustomer(rw http.ResponseWriter, r *http.Request) {

	fmt.Println("Calling to create customer")
}
