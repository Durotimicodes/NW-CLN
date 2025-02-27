package models

import (
	"time"

	"gorm.io/gorm"
)

type Expense struct {
	gorm.Model
	UserID      string    `json:"user_id"`
	Category    string    `json:"category"`
	Amount      float64   `json:"amount"`
	Description string    `json:"description"`
	Date        time.Time `json:"date"`
}
