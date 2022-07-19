package user

import (
	"echo/blog-api/config"
	"echo/blog-api/entity"
	"echo/blog-api/repository/user"
)

type Service struct {
	r    user.UserRepository
	conf config.Config
}

func NewUserService(r user.UserRepository, conf config.Config) UserService {
	return &Service{
		r:    r,
		conf: conf,
	}
}

func (s *Service) FindAllUsers() ([]entity.User, error) {
	users, err := s.r.FindAllUsers()
	return users, err
}

func (s *Service) FindByIDUsers(id uint) (entity.User, error) {
	user, err := s.r.FindByIDUsers(id)
	return user, err
}
