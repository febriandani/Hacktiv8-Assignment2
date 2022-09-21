package orders

import (
	"assign2-new/customers"

	"gorm.io/gorm"
)

type Repository interface {
	CreateOrder(order Order) (Order, error)
	FindByID(id int) (customers.Customer, error)
	GetAllOrders() ([]Order, error)
	Update(order Order) (Order, error)
	GetOrderByID(ID int) (Order, error)
	Delete(order Order) (Order, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db: db}
}

func (r *repository) CreateOrder(order Order) (Order, error) {

	err := r.db.Create(&order).Error
	if err != nil {
		return order, err
	}

	return order, nil

}

func (r *repository) FindByID(id int) (customers.Customer, error) {
	var customer customers.Customer

	// err := r.db.Where("ID = ?", id).Find(&customer).Error
	err := r.db.First(&customer, "id = ?", id).Error
	if err != nil {
		return customer, err
	}

	return customer, nil
}

func (r *repository) GetAllOrders() ([]Order, error) {
	var order []Order

	err := r.db.Find(&order).Error
	if err != nil {
		return order, err
	}

	return order, nil
}

func (r *repository) GetOrderByID(ID int) (Order, error) {
	var order Order
	err := r.db.Find(&order, ID).Error
	if err != nil {
		return order, err
	}

	return order, nil

}

func (r *repository) Update(order Order) (Order, error) {

	// err := r.db.Debug().Update("UPDATE orders SET(item_code,description,quantity,customer_id) VALUES(?,?,?,?) WHERE id = ?", order.ID).Error
	err := r.db.Debug().Save(&order).Error
	if err != nil {
		return order, err
	}
	return order, nil
}

func (r *repository) Delete(order Order) (Order, error) {
	err := r.db.Delete(&order).Error
	if err != nil {
		return order, err
	}

	return order, nil
}
