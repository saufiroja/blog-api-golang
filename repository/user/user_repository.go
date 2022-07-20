package user

import "echo/blog-api/entity"

type UserRepository interface {
	FindAllUsers() ([]entity.User, error)
	FindByIDUsers(id uint) (entity.User, error)
	UpdateUser(id uint, user entity.User) error
	DeleteUser(id uint, user entity.User) error
}
