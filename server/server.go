package server

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/neranjana/samplerest/customer"
	"github.com/neranjana/samplerest/logger"
)

var customerService customer.Service

// RunServer starts an http server with the given port
func RunServer(customerServiceInstance customer.Service, port string) {
	customerService = customerServiceInstance
	logger.Info.Println("Starting server on port", port)
	router := mux.NewRouter()
	router.HandleFunc("/customers", handleGetAllCustomers).Methods("GET")
	router.HandleFunc("/customers/{id}", handleGetSpecificCustomer).Methods("GET")
	router.HandleFunc("/customers", handleStoreCustomer).Methods("POST")
	log.Fatal(http.ListenAndServe("localhost:"+port, router))
}

func handleGetAllCustomers(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(customerService.RetrieveAllCustomers())
}

func handleGetSpecificCustomer(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]
	cust, err := customerService.RetrieveCustomerByID(id)
	if err != nil {
		// json.NewEncoder(w).Encode(nil)
		logger.Info.Println("Could not find customer with ID", id)
	} else {
		json.NewEncoder(w).Encode(cust)
	}

}

func handleStoreCustomer(w http.ResponseWriter, r *http.Request) {
	var newCust customer.Customer
	_ = json.NewDecoder(r.Body).Decode(&newCust)
	logger.Info.Println("New customer", newCust)
	storedCust, err := customerService.StoreCustomer(newCust)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
	} else {
		json.NewEncoder(w).Encode(storedCust)
	}

}
