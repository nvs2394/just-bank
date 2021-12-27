package app

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/nvs2394/just-bank/domain"
	"github.com/nvs2394/just-bank/service"
)

func Start() {
	router := mux.NewRouter()

	customerHandler := CustomerHandlers{
		service: service.NewCustomerService(domain.NewCustomerRepositoryDb()),
	}

	router.HandleFunc("/customers", customerHandler.getCustomers).Methods(http.MethodGet)
	router.HandleFunc("/customers", createCustomer).Methods(http.MethodPost)
	router.HandleFunc("/customers/{customer_id}:[0-9]+", getCustomer).Methods(http.MethodGet)

	log.Fatal(http.ListenAndServe("localhost:8000", router))

}
