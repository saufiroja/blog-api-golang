package user

import "echo/blog-api/entity"

type UserService interface {
	FindAllUsers() ([]entity.User, error)
	FindByIDUsers(id string) (entity.User, error)
	UpdateUser(id string, user entity.User) error
	DeleteUser(id string, user entity.User) error
}
