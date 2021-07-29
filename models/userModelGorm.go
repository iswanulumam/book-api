package models

import "gorm.io/gorm"

type GormUserModel struct {
	db *gorm.DB
}

func (m *GormUserModel) Get() []User {
	var users []User
	m.db.Find(&users)
	return users
}

func NewUserModel(db *gorm.DB) *GormUserModel {
	return &GormUserModel{db: db}
}
