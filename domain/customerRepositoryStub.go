package domain

type CustomerRepositoryStub struct {
	customers []Customer
}

func (s CustomerRepositoryStub) FindAll(status string) ([]Customer, error) {
	return s.customers, nil
}
