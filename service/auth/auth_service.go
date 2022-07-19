package auth

import "echo/blog-api/entity"

type AuthService interface {
	Register(user entity.User) (entity.User, error)
	Login(email, password string) (string, error)
}
