package models

import "gorm.io/gorm"

type Payment struct {
	gorm.Model
	UserID string  `json:"user_id"`
	Amount float64 `json:"amount"`
	Method string  `json:"method"` //Bank Transfer, Card, Direct Debit
	Status string `json:"status"` //Pending, Completed, Failed
}
