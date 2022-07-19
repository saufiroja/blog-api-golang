package user

import (
	"echo/blog-api/entity"

	"gorm.io/gorm"
)

type repository struct {
	DB *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &repository{
		DB: db,
	}
}

func (r *repository) FindAllUsers() ([]entity.User, error) {
	users := []entity.User{}
	err := r.DB.Find(&users).Error
	return users, err
}

func (r *repository) FindByIDUsers(id uint) (entity.User, error) {
	user := entity.User{}
	err := r.DB.Where("id = ?", id).First(&user).Error
	return user, err
}
