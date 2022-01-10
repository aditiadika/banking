package domain

type CustomerRepository interface {
	FindAll() ([]Customer, error)
}

type CustomerRepositoryStub struct {
	customers []Customer
}

func (s CustomerRepositoryStub) FindAll() ([]Customer, error) {
	return s.customers, nil
}

func NewCustomerRepositoryStub() CustomerRepositoryStub {
	customers := []Customer{
		{Id: "1", Name: "John Holt", City: "Kingston", Zipcode: "12345", DateofBirth: "01-01-1980", Status: "active"},
		{Id: "2", Name: "Bruno", City: "San Paulo", Zipcode: "9876", DateofBirth: "01-01-1970", Status: "active"},
		{Id: "3", Name: "Mike", City: "London", Zipcode: "5643", DateofBirth: "01-01-1960", Status: "not active"},
	}
	return CustomerRepositoryStub{customers}
}
