package models

import "gorm.io/gorm"

type GormUserModel struct {
	db *gorm.DB
}

func (m *GormUserModel) Get() ([]User, error) {
	var users []User
	if err := m.db.Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

func (m *GormUserModel) Insert(user User) (User, error) {
	if err := m.db.Save(&user).Error; err != nil {
		return user, err
	}
	return user, nil
}

func NewUserModel(db *gorm.DB) *GormUserModel {
	return &GormUserModel{db: db}
}
