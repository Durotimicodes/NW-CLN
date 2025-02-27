package models

import "time"

type SupportTicket struct {
	ID        string    `json:"id" gorm:"primaryKey"`
	UserID    string    `json:"user_id"`
	Subject   string    `json:"subject"`
	Status    string    `json:"status"` // Open, Closed, Pending
	CreatedAt time.Time `json:"created_at"`
}
