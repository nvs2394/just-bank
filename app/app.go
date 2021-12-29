package app

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/nvs2394/just-bank/domain"
	"github.com/nvs2394/just-bank/service"
)

func sanityCheck() {
	if os.Getenv("SERVER_ADDRESS") == "" || os.Getenv("SERVER_PORT") == "" {
		log.Fatal("Environment variable not defined")
	}
}

func Start() {

	sanityCheck()
	router := mux.NewRouter()

	customerHandler := CustomerHandlers{
		service: service.NewCustomerService(domain.NewCustomerRepositoryDb()),
	}

	router.HandleFunc("/customers", customerHandler.getCustomers).Methods(http.MethodGet)
	router.HandleFunc("/customers", createCustomer).Methods(http.MethodPost)
	router.HandleFunc("/customers/{customer_id}", customerHandler.getCustomer).Methods(http.MethodGet)

	address := os.Getenv("SERVER_ADDRESS")
	port := os.Getenv("SERVER_PORT")

	log.Fatal(http.ListenAndServe(fmt.Sprintf("%s:%s", address, port), router))

}
