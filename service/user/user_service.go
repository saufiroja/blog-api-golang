package user

import "echo/blog-api/entity"

type UserService interface {
	FindAllUsers() ([]entity.User, error)
	FindByIDUsers(id uint) (entity.User, error)
}
