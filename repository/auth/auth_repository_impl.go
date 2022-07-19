package auth

import (
	"echo/blog-api/entity"

	"gorm.io/gorm"
)

type repository struct {
	DB *gorm.DB
}

func NewAuthRepository(db *gorm.DB) AuthRepository {
	return &repository{
		DB: db,
	}
}

func (r *repository) Register(user entity.User) (entity.User, error) {
	err := r.DB.Create(&user).Error
	return user, err
}

func (r *repository) Login(email string) (entity.User, error) {
	user := entity.User{}
	db := r.DB.Where("email = ?", email).First(&user)
	return user, db.Error
}
