package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	FirstName     string    `json:"first_name"`
	LastName      string    `json:"last_name"`
	Email         string    `json:"email" gorm:"unique"`
	Password      string    `json:"password"`
	PhoneNumber   string    `json:"phone_number"`
	DateOfBirth   time.Time `json:"date_of_birth"`
	Address       string    `json:"address"`
	City          string    `json:"city"`
	Country       string    `json:"country"`
	Postcode      string    `json:"postcode"`
	AccountNumber string    `json:"account_number"` // Encrypted before storing
	SortCode      string    `json:"sort_code"`      // Encrypted before storing
	IBAN          string    `json:"iban"`           // Encrypted before storing
	ATMCard       ATMStatus `json:"atm_card"`
	CreditScore   int       `json:"credit_score"`
	Income        float64   `json:"income"`
}

type ATMStatus string

const (
	ATMStatusActive      ATMStatus = "active"
	ATMStatusFrozen      ATMStatus = "frozen"
	ATMStatusDeactivated ATMStatus = "deactivated"
)
