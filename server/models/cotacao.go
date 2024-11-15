package models

import "time"

// Cotacao representa o modelo de cotação no banco de dados.
type Cotacao struct {
	ID        uint      `gorm:"primaryKey"`
	Bid       string    `gorm:"not null"`
	Timestamp time.Time `gorm:"autoCreateTime"`
}
