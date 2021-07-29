package models

type UserModel interface {
	Get() ([]User, error)
	Insert(User) (User, error)
}
