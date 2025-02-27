package models

import "gorm.io/gorm"

type TaxAllocation struct {
	gorm.Model
	UserID string `json:"user_id"`
	Amount float64 `json:"amount"`
	TaxYear string `json:"tax_year"`
}