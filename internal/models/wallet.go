package models

import "time"

type Wallet struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	UserID    uint      `gorm:"not null; uniqueIndex" json:"user_id"`
	Balance   float64   `gorm:"not null" json:"balance"`
	Currency  string    `gorm:"not null" json:"currency"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}