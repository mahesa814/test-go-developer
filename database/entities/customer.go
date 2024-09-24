package entities

import (
	"gorm.io/gorm"
	"time"
)

type Customer struct {
	gorm.Model
	ID           string     `gorm:"type:char(36);primaryKey;default:UUID();"`
	NIK          string     `gorm:"column:nik;type:varchar(255);not null"`        // Specify length for varchar
	FullName     string     `gorm:"column:full_name;type:varchar(255);not null"`  // Specify length for varchar
	LegalName    string     `gorm:"column:legal_name;type:varchar(255);not null"` // Specify length for varchar
	DateOfBirth  time.Time  `gorm:"column:date_of_birth;type:date;"`
	PlaceOfBirth string     `gorm:"column:place_of_birth;type:text;"`       // Change type to text
	Salary       int64      `gorm:"column:salary;type:int;default:0;"`      // Use type:int for integer
	KtpPhoto     string     `gorm:"column:ktp_photo;type:varchar(255);"`    // Specify length for varchar
	SelfiePhoto  string     `gorm:"column:selfie_photo;type:varchar(255);"` // Specify length for varchar
	CreatedAt    *time.Time `json:"created_at" gorm:"column:created_at;default:CURRENT_TIMESTAMP;"`
	UpdatedAt    *time.Time `json:"updated_at" gorm:"column:updated_at;default:CURRENT_TIMESTAMP;"`
}
