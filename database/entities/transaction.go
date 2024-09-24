package entities

type Transaction struct {
	ID          string   `gorm:"column:id;primary_key;type:uuid;default:uuid_generate_v4();not null"`
	PhoneNumber string   `gorm:"column:phone_number;primary_key;type:varchar(15);not null"`
	OTR         float64  `gorm:"column:otr;type:integer"`
	AdminFee    float64  `gorm:"column:admin_fee;type:decimal(10,2);not null"`
	Installment float64  `gorm:"column:installment;type:decimal(10,2);not null"`
	Interest    float64  `gorm:"column:interest;type:decimal(10,2);not null"`
	AssetName   string   `gorm:"column:asset_name;type:varchar(255);not null"`
	CustomerID  string   `gorm:"column:customer_id;type:varchar(255);not null"`
	Customer    Customer `gorm:"foreignkey:CustomerID;references:ID"`
}
