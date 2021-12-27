package domain

type CustomerRepositoryStub struct {
	customers []Customer
}

func (s CustomerRepositoryStub) FindAll() ([]Customer, error) {
	return s.customers, nil
}

func NewCustomerRepositoryStub() CustomerRepositoryStub {
	customers := []Customer{
		{Id: "1001", Name: "Sonny", City: "Singapore", Zipcode: "115000", DateOfBirth: "09/08/1994", Status: "ACTIVE"},
		{Id: "1001", Name: "Bao Bao", City: "Viet Nam", Zipcode: "115000", DateOfBirth: "09/08/1996", Status: "ACTIVE"},
	}

	return CustomerRepositoryStub{customers: customers}
}
