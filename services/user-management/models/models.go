package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	FirstName   string    `json:"first_name"`
	LastName    string    `json:"last_name"`
	Email       string    `json:"email" gorm:"unique"`
	Password    string    `json:"password"`
	PhoneNumber string    `json:"phone_number"`
	DateOfBirth time.Time `json:"date_of_birth"`
	Address     string    `json:"address"`
	City        string    `json:"city"`
	Country     string    `json:"country"`
	Postcode    string    `json:"postcode"`
	Occupation  string    `json:"occupation"`
	CreditScore int       `json:"credit_score"`
	Income      float64   `json:"income"`
	Company     string    `json:"company"`
}
