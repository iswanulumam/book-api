package models

import (
	"alta/book-api/api/middlewares"

	"gorm.io/gorm"
)

// Model Customer

type Customer struct {
	gorm.Model
	Name     string
	Email    string
	Gender   string `sql:"type:ENUM('male', 'female')"`
	Password string `gorm:"<-:false"`
	Token    string `gorm:"<-:false"`
}

// type Customer struct {
// 	gorm.Model
// 	Name     string `json:"name" form:"name"`
// 	Email    string `json:"email" form:"email"`
// 	Password string `json:"password" form:"password"`
// 	Token    string `json:"token" form:"token"`
// }

type GormCustomerModel struct {
	db *gorm.DB
}

func NewCustomerModel(db *gorm.DB) *GormCustomerModel {
	return &GormCustomerModel{db: db}
}

// Interface Customer

type CustomerModel interface {
	GetAll() ([]Customer, error)
	Get(customerId int) (Customer, error)
	Insert(Customer) (Customer, error)
	Edit(customer Customer, customerId int) (Customer, error)
	Delete(customerId int) (Customer, error)
	Login(email, password string) (Customer, error)
}

func (m *GormCustomerModel) GetAll() ([]Customer, error) {
	var customer []Customer
	if err := m.db.Find(&customer).Error; err != nil {
		return nil, err
	}
	return customer, nil
}

func (m *GormCustomerModel) Get(customerId int) (Customer, error) {
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

func (m *GormCustomerModel) Edit(newCustomer Customer, customerId int) (Customer, error) {
	var customer Customer
	if err := m.db.Find(&customer, "id=?", customerId).Error; err != nil {
		return customer, err
	}

	customer.Name = newCustomer.Name
	customer.Email = newCustomer.Email
	customer.Password = newCustomer.Password

	if err := m.db.Save(&customer).Error; err != nil {
		return customer, err
	}
	return customer, nil
}

func (m *GormCustomerModel) Delete(customerId int) (Customer, error) {
	var customer Customer
	if err := m.db.Find(&customer, "id=?", customerId).Error; err != nil {
		return customer, err
	}
	if err := m.db.Delete(&customer).Error; err != nil {
		return customer, err
	}
	return customer, nil
}

func (m *GormCustomerModel) Login(email, password string) (Customer, error) {
	var customer Customer
	var err error

	if err = m.db.Where("email = ? AND password = ?", email, password).First(&customer).Error; err != nil {
		return customer, err
	}

	customer.Token, err = middlewares.CreateToken(int(customer.ID))

	if err != nil {
		return customer, err
	}

	if err := m.db.Save(customer).Error; err != nil {
		return customer, err
	}

	return customer, nil
}
