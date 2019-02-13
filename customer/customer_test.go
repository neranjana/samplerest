package customer

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRetrieveAllCustomers(t *testing.T) {

	// Given
	var customers []Customer
	customers = append(customers, Customer{ID: "1", Firstname: "John", Lastname: "Smith"})
	customers = append(customers, Customer{ID: "2", Firstname: "Sam", Lastname: "Ford"})

	simpleCustomerService := SimpleService{customers}

	// When
	allCustomers := simpleCustomerService.RetrieveAllCustomers()

	// Then
	// test whether retrieves two Customer instances
	assert.Equal(t, 2, len(allCustomers))

	// test the 0th Customer's data values
	assert.Equal(t, "1", allCustomers[0].ID)
	assert.Equal(t, "John", allCustomers[0].Firstname)
	assert.Equal(t, "Smith", allCustomers[0].Lastname)

	// test the second customer
	assert.Equal(t, Customer{ID: "2", Firstname: "Sam", Lastname: "Ford"}, allCustomers[1])
}

func TestRetrieveCustomerByIDCustomerNotFound(t *testing.T) {
	//Given
	var customers []Customer
	customers = append(customers, Customer{ID: "1", Firstname: "John", Lastname: "Smith"})
	customers = append(customers, Customer{ID: "2", Firstname: "Sam", Lastname: "Ford"})

	simpleCustomerService := SimpleService{customers}

	//When
	customer, err := simpleCustomerService.RetrieveCustomerByID("99")

	//Then
	if assert.Error(t, err) {
		// error is the correct error
		expectedError := fmt.Errorf("Customer with id %s not found", "99")
		assert.Equal(t, expectedError, err)
	}

	// customer is an empty cutomer
	assert.Equal(t, Customer{}, customer)
}

func TestRetrieveCustomerByIDCustomerFound(t *testing.T) {
	//Given
	var customers []Customer
	customers = append(customers, Customer{ID: "1", Firstname: "John", Lastname: "Smith"})
	customers = append(customers, Customer{ID: "2", Firstname: "Sam", Lastname: "Ford"})
	customers = append(customers, Customer{ID: "3", Firstname: "Joseph", Lastname: "Biden"})

	simpleCustomerService := SimpleService{customers}

	//When
	customer, err := simpleCustomerService.RetrieveCustomerByID("2")

	//Then
	// error is nil
	assert.Nil(t, err)
	// Customer is the correct customer
	assert.Equal(t, Customer{ID: "2", Firstname: "Sam", Lastname: "Ford"}, customer)

}

func TestStoreCustomerWithNonExistingID(t *testing.T) {
	//Given
	var customers []Customer
	customers = append(customers, Customer{ID: "1", Firstname: "John", Lastname: "Smith"})
	customers = append(customers, Customer{ID: "2", Firstname: "Sam", Lastname: "Ford"})
	customers = append(customers, Customer{ID: "3", Firstname: "Joseph", Lastname: "Biden"})

	simpleCustomerService := SimpleService{customers}

	newCustomer := Customer{ID: "4", Firstname: "James", Lastname: "Bond"}

	//When
	customerID, err := simpleCustomerService.StoreCustomer(newCustomer)

	//Then
	assert.Nil(t, err)
	assert.Equal(t, "4", customerID)
	assert.Equal(t, 4, len(simpleCustomerService.customers))
	assert.Equal(t, "James", simpleCustomerService.customers[3].Firstname)
	assert.Equal(t, "Bond", simpleCustomerService.customers[3].Lastname)

}

func TestStoreCustomerWithExistingID(t *testing.T) {
	//Given
	var customers []Customer
	customers = append(customers, Customer{ID: "1", Firstname: "John", Lastname: "Smith"})
	customers = append(customers, Customer{ID: "2", Firstname: "Sam", Lastname: "Ford"})
	customers = append(customers, Customer{ID: "3", Firstname: "Joseph", Lastname: "Biden"})

	simpleCustomerService := SimpleService{customers}

	newCustomer := Customer{ID: "3", Firstname: "James", Lastname: "Bond"}

	//When
	customerID, err := simpleCustomerService.StoreCustomer(newCustomer)

	//Then
	assert.Equal(t, fmt.Sprintf("Customer with id %s already exists", "3"), err.Error())
	assert.Equal(t, "", customerID)
	assert.Equal(t, 3, len(simpleCustomerService.customers))

}

func TestRemoveCustomerThatExists(t *testing.T) {
	//Given
	var customers []Customer
	customers = append(customers, Customer{ID: "1", Firstname: "John", Lastname: "Smith"})
	customers = append(customers, Customer{ID: "2", Firstname: "Sam", Lastname: "Ford"})
	customers = append(customers, Customer{ID: "3", Firstname: "Joseph", Lastname: "Biden"})

	simpleCustomerService := SimpleService{customers}

	//When
	customerID, err := simpleCustomerService.RemoveCustomer("2")
	//Then
	assert.Nil(t, err)
	assert.Equal(t, "2", customerID)
	assert.Equal(t, 2, len(simpleCustomerService.customers))
}

func TestRemoveCustomerThatDoesNotExist(t *testing.T) {
	//Given
	var customers []Customer
	customers = append(customers, Customer{ID: "1", Firstname: "John", Lastname: "Smith"})
	customers = append(customers, Customer{ID: "2", Firstname: "Sam", Lastname: "Ford"})
	customers = append(customers, Customer{ID: "3", Firstname: "Joseph", Lastname: "Biden"})

	simpleCustomerService := SimpleService{customers}

	//When
	customerID, err := simpleCustomerService.RemoveCustomer("99")
	//Then
	assert.Equal(t, fmt.Sprintf("Customer with id %s does not exist", "99"), err.Error())
	assert.Equal(t, "", customerID)
	assert.Equal(t, 3, len(simpleCustomerService.customers))
}
