package user

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name     string `json:"name"  gorm:"type:varchar(100);not null"`
	Email    string `json:"email"  gorm:"type:varchar(300);not null;uniqueIndex"`
	Phone    string `json:"phone" gorm:"type:varchar(11);not null;uniqueIndex"`
	Password string `json:"password" gorm:"type:varchar(255);not null"`
}