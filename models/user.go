package models

type UserModel interface {
	Get() ([]User, error)
	Insert(User) (User, error)
	// NewUserModel(db *gorm.DB) *GormUserModel
}
