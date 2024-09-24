package entities

import (
	"gorm.io/gorm"
	"time"
)

type LoanLimit struct {
	gorm.Model
	ID         string     `gorm:"type:char(36);primaryKey;default:UUID();"`
	Tenor      int        `gorm:"column:tenor;type:int;not null"`            // Use type:int for integer
	Limit      int64      `gorm:"column:limit;type:int;default:0;not null"`  // Use type:int for integer
	CustomerID string     `gorm:"column:customer_id;type:char(36);not null"` // Change type to char(36)
	Customer   Customer   `gorm:"foreignKey:CustomerID;references:ID"`       // Corrected foreign key tag
	CreatedAt  *time.Time `json:"created_at" gorm:"column:created_at;default:CURRENT_TIMESTAMP;"`
	UpdatedAt  *time.Time `json:"updated_at" gorm:"column:updated_at;default:CURRENT_TIMESTAMP;"`
}
