package entities

type LoanLimit struct {
	ID         string   `gorm:"column:id;primary_key;type:uuid;default:uuid_generate_v4();not null"`
	Tenor      int      `gorm:"column:tenor;type:integer;not null"`
	Limit      int64    `gorm:"column:limit;type:integer;default:0;not null"`
	CustomerID string   `gorm:"column:customer_id;type:uuid;not null"`
	Customer   Customer `gorm:"foreignkey:CustomerID;references:ID"`
}
