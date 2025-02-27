package models

import "gorm.io/gorm"

type Transaction struct {
	gorm.Model
	UserID string  `json:"user_id"`
	Type   string  `json:"type"` //Credit, Debit
	Amount float64 `json:"amount"`
	Source string  `json:"source"`
}
