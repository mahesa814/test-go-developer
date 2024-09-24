package entities

import (
	"gorm.io/gorm"
	"time"
)

type Customer struct {
	gorm.Model
	ID           string      `json:"id" gorm:"type:char(36);primaryKey;default:UUID();"`
	NIK          string      `json:"nik" gorm:"column:nik;type:varchar(255);not null"`               // Specify length for varchar
	FullName     string      `json:"full_name" gorm:"column:full_name;type:varchar(255);not null"`   // Specify length for varchar
	LegalName    string      `json:"legal_name" gorm:"column:legal_name;type:varchar(255);not null"` // Specify length for varchar
	DateOfBirth  time.Time   `json:"date_of_birth" gorm:"column:date_of_birth;type:date;"`
	PlaceOfBirth string      `json:"place_of_birth" gorm:"column:place_of_birth;type:text;"`     // Change type to text
	Salary       int64       `json:"salary" gorm:"column:salary;type:int;default:0;"`            // Use type:int for integer
	KtpPhoto     string      `json:"ktp_photo" gorm:"column:ktp_photo;type:varchar(255);"`       // Specify length for varchar
	SelfiePhoto  string      `json:"selfie_photo" gorm:"column:selfie_photo;type:varchar(255);"` // Specify length for varchar
	LoanLimit    []LoanLimit `gorm:"foreignKey:CustomerID;references:ID"`                        // Corrected foreign key reference
	CreatedAt    *time.Time  `json:"created_at" gorm:"column:created_at;default:CURRENT_TIMESTAMP;"`
	UpdatedAt    *time.Time  `json:"updated_at" gorm:"column:updated_at;default:CURRENT_TIMESTAMP;"`
}
