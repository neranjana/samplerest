package customer

import (
	"errors"
	"fmt"
)

// Customer is a struct that represents a customer
type Customer struct {
	ID        string `json:"id,omitempty"`
	Firstname string `json:"firstname,omitempty"`
	Lastname  string `json:"lastname,omitempty"`
}

// Service that handles Customer CRUD operations
type Service interface {
	StoreCustomer(newCustomer Customer) (string, error)
	RetrieveAllCustomers() []Customer
	RetrieveCustomerByID(ID string) (Customer, error)
}

// SimpleService that handles Customer CRUD operaitons using an in memory slice
type SimpleService struct {
	customers []Customer
}

// RetrieveAllCustomers returns all the Customers
func (simpleService *SimpleService) RetrieveAllCustomers() []Customer {
	return simpleService.customers
}

// RetrieveCustomerByID returns a specific customer with the given ID
func (simpleService *SimpleService) RetrieveCustomerByID(ID string) (Customer, error) {
	var customer Customer
	var customerError error
	customerFound := false

	for _, cust := range simpleService.customers {
		if cust.ID == ID {
			customer = cust
			customerFound = true
		}
	}

	if !customerFound {
		customerError = errors.New("Customer with id " + ID + " not found")
	}
	// customer = Customer{ID: "1", Firstname: "John", Lastname: "Smith"}
	return customer, customerError
}

// StoreCustomer adds the customer to the storage and returns the ID of the customer
func (simpleService *SimpleService) StoreCustomer(newCustomer Customer) (string, error) {
	var returnCustID string
	var returnError error
	_, err := simpleService.RetrieveCustomerByID(newCustomer.ID)
	if err != nil && err.Error() == fmt.Sprintf("Customer with id %s not found", newCustomer.ID) {
		simpleService.customers = append(simpleService.customers, newCustomer)
		returnCustID = newCustomer.ID
		returnError = nil
	} else if err == nil {
		// returnCustID = nil
		returnError = fmt.Errorf("Customer with id %s already exists", newCustomer.ID)
	} else {
		// returnCustID = nil
		returnError = fmt.Errorf("Something went wrong %s", err.Error())
	}

	return returnCustID, returnError
}

// RemoveCustomer removes the customer with the given ID from the storage and returns the ID
func (simpleService *SimpleService) RemoveCustomer(ID string) (string, error) {
	var returnCustID string
	var returnError error
	var indexOfCustomer int
	customerFound := false

	for currentIndex, cust := range simpleService.customers {
		if cust.ID == ID {
			indexOfCustomer = currentIndex
			customerFound = true
		}
	}

	if customerFound {
		simpleService.customers = simpleService.customers[:indexOfCustomer+copy(simpleService.customers[indexOfCustomer:], simpleService.customers[indexOfCustomer+1:])]
		returnCustID = ID
	} else {
		returnError = fmt.Errorf("Customer with id %s does not exist", ID)
	}

	return returnCustID, returnError

}
