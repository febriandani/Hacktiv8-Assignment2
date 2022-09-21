package customers

import "gorm.io/gorm"

type Repository interface {
	Save(customer Customer) (Customer, error)
	FindByID(id int) (Customer, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) Save(customer Customer) (Customer, error) {
	err := r.db.Debug().Create(&customer).Error
	if err != nil {
		return customer, err
	}

	return customer, nil
}

func (r *repository) FindByID(id int) (Customer, error) {
	var customer Customer

	err := r.db.Where("ID = ?", id).Find(&customer).Error
	if err != nil {
		return customer, err
	}

	return customer, nil
}
