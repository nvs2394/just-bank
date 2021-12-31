package app

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
	"github.com/nvs2394/just-bank/domain"
	"github.com/nvs2394/just-bank/service"
)

func sanityCheck() {
	if os.Getenv("SERVER_ADDRESS") == "" || os.Getenv("SERVER_PORT") == "" {
		log.Fatal("Environment variable not defined")
	}
}

func getDBClient() *sqlx.DB {
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
	return client
}

func Start() {

	sanityCheck()
	router := mux.NewRouter()
	dbClient := getDBClient()

	customerRepositoryDB := domain.NewCustomerRepositoryDb(dbClient)
	accountRepositoryDB := domain.NewAccountRepositoryDb(dbClient)
	transactionRepositoryDB := domain.NewTransactionRepositoryDb(dbClient)

	customerHandler := CustomerHandlers{
		service: service.NewCustomerService(customerRepositoryDB),
	}

	accountHandler := AccountHandler{
		service: service.NewAccountService(accountRepositoryDB),
	}

	transactionHandler := TransactionHandler{
		service: service.NewTransactionService(transactionRepositoryDB),
	}

	// Customer routes
	router.HandleFunc("/customers", customerHandler.getCustomers).Methods(http.MethodGet)
	router.HandleFunc("/customers", createCustomer).Methods(http.MethodPost)
	router.HandleFunc("/customers/{customer_id}", customerHandler.getCustomer).Methods(http.MethodGet)
	router.HandleFunc("/customers/{customer_id}/accounts", accountHandler.CreateAccount).Methods(http.MethodPost)

	// Transaction routes"
	router.HandleFunc("/transactions", transactionHandler.CreateTransaction).Methods(http.MethodPost)

	address := os.Getenv("SERVER_ADDRESS")
	port := os.Getenv("SERVER_PORT")

	log.Fatal(http.ListenAndServe(fmt.Sprintf("%s:%s", address, port), router))

}
