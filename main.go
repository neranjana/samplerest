package main

import (
	"io/ioutil"
	"os"

	"github.com/neranjana/samplerest/customer"
	"github.com/neranjana/samplerest/logger"
	"github.com/neranjana/samplerest/server"
)

func main() {

	logger.Init(ioutil.Discard, os.Stdout, os.Stdout, os.Stderr)

	// Pointer to customer.SimpleService
	simpleCustomerService := new(customer.SimpleService)
	initSampleCustomers(simpleCustomerService)

	port := "8902"
	server.RunServer(simpleCustomerService, port)
}

func initSampleCustomers(simpleService *customer.SimpleService) {
	simpleService.StoreCustomer(customer.Customer{ID: "1", Firstname: "John", Lastname: "Smith"})
	simpleService.StoreCustomer(customer.Customer{ID: "2", Firstname: "Sam", Lastname: "Ford"})

}
