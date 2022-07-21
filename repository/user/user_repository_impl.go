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
	if err != nil {
		return users, err
	}
	return users, err
}

func (r *repository) FindByIDUsers(id string) (entity.User, error) {
	user := entity.User{}
	err := r.DB.Where("id = ?", id).First(&user).Error
	if err != nil {
		return user, err
	}
	return user, err
}

func (r *repository) UpdateUser(id string, user entity.User) error {
	err := r.DB.Model(&user).Where("id = ?", id).Updates(&user).Error
	if err != nil {
		return err
	}
	return err
}

func (r *repository) DeleteUser(id string, user entity.User) error {
	err := r.DB.Where("id = ?", id).Delete(&user).Error
	if err != nil {
		return err
	}
	return err
}
