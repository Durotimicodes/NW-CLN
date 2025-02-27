package models

import (
	"time"

	"gorm.io/gorm"
)

type BNPL struct {
	gorm.Model
	Amount       float64 `json:"amount"`
	InterestRate float64 `json:"interest_rate"`
	Installments int     `json:"installments"`
	Status       string  `json:"status"`
}


type Installment struct {
	gorm.Model
	BNPLID string `json:"bnpl_id"`
	Amount float64 `json:"amount"`
	DueDate time.Time `json:"due_date"`
	Status string `json:"status"` //Pending, Paid, Overdue
}