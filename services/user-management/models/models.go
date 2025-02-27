package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	UserName string `json:"user_name"`
	Email string `json:"email" gorm:"unique"`
	Password string `json:"password"`
}