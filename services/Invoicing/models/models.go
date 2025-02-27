package models

import (
	"time"

	"gorm.io/gorm"
)

// An instance of an invoice
type Invoice struct {
	gorm.Model
	UserID  string    `json:"user_id"`
	Amount  float64   `json:"amount"`
	Status  string    `json:"status"` //Pending , Paid , Overdue
	DueDate time.Time `json:"due_date"`
}
