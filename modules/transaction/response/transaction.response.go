package response

import (
	responseCustomer "test-go-developer/modules/customer/response"
	"time"
)

type TransactionsResponse struct {
	ID             string                              `json:"id"`
	Customer       *responseCustomer.CustomersResponse `json:"customer"`
	ContractNumber string                              `json:"contract_number"`
	OTR            float64                             `json:"otr"` // Changed to decimal for currency
	AdminFee       float64                             `json:"admin_fee"`
	Installment    float64                             `json:"installment"`
	InterestAmount float64                             `json:"interest_amount"`
	AssetName      string                              `json:"asset_name"`
	CreatedAt      *time.Time                          `json:"created_at"`
	UpdatedAt      *time.Time                          `json:"updated_at"`
}
