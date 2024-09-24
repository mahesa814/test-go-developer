package entities

import (
	"gorm.io/gorm"
	"time"
)

type Transaction struct {
	gorm.Model
	ID             string     `gorm:"type:char(36);primaryKey;default:UUID();"`
	ContractNumber string     `gorm:"column:contract_number;type:varchar(15);not null;unique"`
	OTR            float64    `gorm:"column:otr;type:decimal(10,2);not null"` // Changed to decimal for currency
	AdminFee       float64    `gorm:"column:admin_fee;type:decimal(10,2);not null"`
	Installment    float64    `gorm:"column:installment;type:decimal(10,2);not null"`
	InterestAmount float64    `gorm:"column:interest_amount;type:decimal(10,2);not null"`
	AssetName      string     `gorm:"column:asset_name;type:varchar(255);not null"`
	CustomerID     string     `gorm:"column:customer_id;type:char(36);not null"` // Ensure it matches the UUID type
	Customer       Customer   `gorm:"foreignkey:CustomerID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	LoanLimitId    string     `gorm:"column:loan_limit_id;type:char(36);not null"` // Ensure it matches the UUID type
	LoanLimit      LoanLimit  `gorm:"foreignkey:LoanLimitId;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	CreatedAt      *time.Time `json:"created_at" gorm:"column:created_at;default:CURRENT_TIMESTAMP;"`
	UpdatedAt      *time.Time `json:"updated_at" gorm:"column:updated_at;default:CURRENT_TIMESTAMP;"`
}
