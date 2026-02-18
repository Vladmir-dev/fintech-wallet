package models
import "time"

type Transaction struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	WalletID  uint      `gorm:"not null" json:"wallet_id"`
	Amount    float64   `gorm:"not null" json:"amount"`
	Type      string    `gorm:"not null" json:"type"` // "credit" or "debit"
	Reference string    `gorm:"not null" json:"reference"`
	Counterparty string    `gorm:"not null" json:"counterparty"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}