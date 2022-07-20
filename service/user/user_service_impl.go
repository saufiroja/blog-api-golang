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
	return s.r.FindAllUsers()
}

func (s *Service) FindByIDUsers(id uint) (entity.User, error) {
	return s.r.FindByIDUsers(id)
}

func (s *Service) UpdateUser(id uint, user entity.User) error {
	// check if id is exist
	_, err := s.r.FindByIDUsers(id)
	if err != nil {
		return err
	}
	// update user
	return s.r.UpdateUser(id, user)
}

func (s *Service) DeleteUser(id uint, user entity.User) error {
	// check if id is exist
	_, err := s.r.FindByIDUsers(id)
	if err != nil {
		return err
	}
	// delete user
	return s.r.DeleteUser(id, user)
}
