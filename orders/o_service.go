package orders

type Service interface {
	Create(input OrderInput) (Order, error)
	GetOrders() ([]Order, error)
	UpdateOrder(inputID GetOrderByID, inputData OrderUpdate) (Order, error)
	GetOrderByID(inputID GetOrderByID) (Order, error)
	DeleteOrder(inputID int) (Order, error)
}

type service struct {
	r Repository
}

func NewService(r Repository) *service {
	return &service{r}
}

func (s *service) Create(input OrderInput) (Order, error) {
	orderCreate := Order{}

	orderCreate.Customer_ID = input.Customer_id
	orderCreate.Item_code = input.Items_code
	orderCreate.Description = input.Description
	orderCreate.Quantity = input.Quantity

	//check apakah customer id ada atau tidak di db jika ada maka lanjut proses input
	check, err := s.r.FindByID(orderCreate.Customer_ID)
	if check.ID == 0 {
		return orderCreate, err
	}

	newOrder, err := s.r.CreateOrder(orderCreate)
	if err != nil {
		return newOrder, err
	}

	return newOrder, nil

}

func (s *service) GetOrders() ([]Order, error) {

	orders, err := s.r.GetAllOrders()
	if err != nil {
		return orders, err
	}

	return orders, nil
}

func (s *service) GetOrderByID(inputID GetOrderByID) (Order, error) {
	order, err := s.r.GetOrderByID(inputID.Order_ID)
	if err != nil {
		return order, err
	}

	return order, nil
}

func (s *service) UpdateOrder(inputID GetOrderByID, inputData OrderUpdate) (Order, error) {

	order, err := s.r.GetOrderByID(inputID.Order_ID)
	if err != nil {
		return order, err
	}

	order.Item_code = inputData.Items_code
	order.Description = inputData.Description
	order.Quantity = inputData.Quantity
	order.Customer_ID = inputData.Customer_id

	updateOrder, err := s.r.Update(order)
	if err != nil {
		return updateOrder, err
	}

	return updateOrder, nil

}

func (s *service) DeleteOrder(inputID int) (Order, error) {
	order, err := s.r.GetOrderByID(inputID)
	if err != nil {
		return order, err
	}

	deleteOrder, err := s.r.Delete(order)
	if err != nil {
		return deleteOrder, err
	}
	return deleteOrder, nil
}
