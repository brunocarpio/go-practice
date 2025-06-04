package main

type CustomerService struct {
	customerRepository interface {
		getCustomer(id string) (string, error)
	}
}

type InMemoryCustomerRepository struct {
	movies map[string]string
}

func (cr InMemoryCustomerRepository) getCustomer(id string) (string, error) {
	return cr.movies[id], nil
}
