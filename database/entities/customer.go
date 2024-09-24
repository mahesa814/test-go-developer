package entities

import "time"

type Customer struct {
	ID           string    `gorm:"column:id;primary_key;type:uuid;default:uuid_generate_v4();not null"`
	NIK          string    `gorm:"column:nik;type:varchar;not null"`
	FullName     string    `gorm:"column:full_name;type:varchar;not null"`
	LegalName    string    `gorm:"column:legal_name;type:varchar;not null"`
	DateOfBirth  time.Time `gorm:"column:date_of_birth;type:date;"`
	PlaceOfBirth string    `gorm:"column:place_of_birth;type:string;"`
	Salary       int64     `gorm:"column:salary;type:integer; default:0;"`
	KtpPhoto     string    `gorm:"column:ktp_photo;type:varchar;"`
	SelfiePhoto  string    `gorm:"column:selfie_photo;type:varchar;"`
}
