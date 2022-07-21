package entity

import "time"

type Article struct {
	ID          string `json:"id" gorm:"primary_key;auto_increment"`
	Title       string `json:"title" validate:"required"`
	Description string `json:"description" validate:"required"`

	UserID string `json:"user_id" gorm:"not null"`
	User   User   `json:"user" gorm:"foreignkey:UserID;association_foreignkey:ID;->:false;<-:create"`

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	DeletedAt time.Time `json:"deleted_at"`
}
