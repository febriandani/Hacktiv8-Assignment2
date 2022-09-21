package customers

type Service interface {
	CreateCustomer(input CustomerInput) (Customer, error)
}

type service struct {
	r Repository
}

func NewCustomerService(r Repository) *service {
	return &service{r}
}

func (s *service) CreateCustomer(input CustomerInput) (Customer, error) {
	customerInput := Customer{}

	customerInput.Customer_Name = input.Customer_Name

	newCustomer, err := s.r.Save(customerInput)
	if err != nil {
		return newCustomer, err
	}

	return newCustomer, nil
}
