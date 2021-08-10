package models

type UserModel interface {
	Get() ([]User, error)
	GetOne(userId int) (User, error)
	Insert(User) (User, error)
}
