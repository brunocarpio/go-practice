package main

import "fmt"

func main() {
	customerRepository := InMemoryCustomerRepository{
		movies: map[string]string{"Holly": "CF-1223"}}

	customerService := CustomerService{
		customerRepository,
	}

	customerId, _ := customerService.customerRepository.getCustomer("Holly")
	fmt.Printf("customerId: %v\n", customerId)
}
