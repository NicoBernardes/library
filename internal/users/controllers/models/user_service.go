package models

type UserService interface {
	CreateUser(user *User) error
	GetUser(id int64) (*User, error)
	GetAllUsers() ([]*User, error)
	UpdateUser(id int64, user *User) error
	DeleteUser(id int64) error
}
