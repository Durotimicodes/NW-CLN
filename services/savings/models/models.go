package models

import "gorm.io/gorm"

type Savings struct {
	gorm.Model
	UserID   string  `json:"user_id"`
	Balance  float64 `json:"balance"`
	Interest float64 `json:"interest"`
}
