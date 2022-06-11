package app

import (
	"banking/domain"
	"banking/service"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

func sanityCheck() {
	if os.Getenv("SERVER_ADDRESS") == "" ||
		os.Getenv("SERVER_PORT") == "" {
		log.Fatal("Enviroment variable not defined...")
	}
}

func Start() {
	router := mux.NewRouter()

	sanityCheck()

	//wiring
	ch := CustomerHandlers{service: service.NewCustomerService(domain.NewCustomerRepositoryDb())}
	// define routes
	router.HandleFunc("/customers", ch.getAllCustomers).Methods(http.MethodGet)
	router.HandleFunc("/customers/{customer_id:[0-9]+}", ch.getCustomer).Methods(http.MethodGet)

	// starting server
	address := os.Getenv("SERVER_ADDRESS")
	port := os.Getenv("SERVER_PORT")
	log.Fatal(http.ListenAndServe(fmt.Sprintf("%s:%s", address, port), router))
}
