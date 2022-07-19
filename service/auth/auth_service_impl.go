package auth

import (
	"echo/blog-api/config"
	"echo/blog-api/entity"
	"echo/blog-api/repository/auth"
	"echo/blog-api/utils"

	"golang.org/x/crypto/bcrypt"
)

type Service struct {
	r    auth.AuthRepository
	conf config.Config
}

func NewAuthService(r auth.AuthRepository, conf config.Config) AuthService {
	return &Service{
		r:    r,
		conf: conf,
	}
}

func (s *Service) Register(user entity.User) (entity.User, error) {
	// hash password
	hash, _ := bcrypt.GenerateFromPassword([]byte(user.Password), 12)
	user.Password = string(hash)
	// return user
	return s.r.Register(user)
}

func (s *Service) Login(email, password string) (string, error) {
	// compare password
	user, _ := s.r.Login(email)
	bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	// generate token
	token, _ := utils.CreateToken(user.ID, s.conf.JWT_SECRET)

	return token, nil
}
