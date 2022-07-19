package entity

import (
	"time"
)

// User is the user entity.
type User struct {
	ID       uint   `gorm:"primaryKey;autoIncrement;not null" json:"id"`
	Username string `gorm:"size:50;not null;unique" json:"username" validate:"required,min=4,max=50"`
	Email    string `gorm:"size:100;not null;unique" json:"email" validate:"required"`
	Password string `gorm:"size:100;not null" json:"password" validate:"required,min=8,max=100"`

	CreatedAt time.Time  `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt time.Time  `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at"`
}

// // Before create hash the password.
// func (u *User) BeforeCreate(tx *gorm.DB) error {
// 	salt := 10
// 	password := []byte(u.Password)
// 	hash, _ := bcrypt.GenerateFromPassword(password, salt)
// 	u.Password = string(hash)
// 	return nil
// }
