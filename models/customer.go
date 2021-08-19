package models

import "gorm.io/gorm"

// Model Customer

type Customer struct {
	gorm.Model
	Name     string `json:"name" form:"name"`
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
	Token    string `json:"token" form:"token"`
}

// type APICustomer struct {
// 	gorm.Model
// 	Name     string `json:"name" form:"name"`
// 	Email    string `json:"email" form:"email"`
// }

type GormCustomerModel struct {
	db *gorm.DB
}

func NewCustomerModel(db *gorm.DB) *GormCustomerModel {
	return &GormCustomerModel{db: db}
}

// Interface Customer

type CustomerModel interface {
	Get() ([]Customer, error)
	GetOne(customerId int) (Customer, error)
	Insert(Customer) (Customer, error)
}

func (m *GormCustomerModel) Get() ([]Customer, error) {
	var customer []Customer
	if err := m.db.Find(&customer).Error; err != nil {
		return nil, err
	}
	return customer, nil
}

func (m *GormCustomerModel) GetOne(customerId int) (Customer, error) {
	var customer Customer
	if err := m.db.Find(&customer, customerId).Error; err != nil {
		return customer, err
	}
	return customer, nil
}

func (m *GormCustomerModel) Insert(customer Customer) (Customer, error) {
	if err := m.db.Save(&customer).Error; err != nil {
		return customer, err
	}
	return customer, nil
}
