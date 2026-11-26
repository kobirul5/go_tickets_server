package user

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name     string `json:"name" validate:"required" gorm:"type:varchar(100);not null"`
	Email    string `json:"email" validate:"required,email" gorm:"type:varchar(300);not null;uniqueIndex"`
	Phone    string `json:"phone" validate:"required,min=11" gorm:"type:varchar(11);not null;uniqueIndex"`
	Password string `json:"password" validate:"required,min=8" gorm:"type:varchar(255);not null"`
}