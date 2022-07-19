package auth

import "echo/blog-api/entity"

type AuthRepository interface {
	Register(user entity.User) (entity.User, error)
	Login(email string) (entity.User, error)
}
