package request

type TransactionRequest struct {
	ContactNumber        string `json:"contact_number"`
	OTR                  string `json:"otr"`
	AdminFee             string `json:"admin_fee"`
	AmountOfInstallments string `json:"amount_of_installments"`
	AmountOfInterest     string `json:"amount_of_interest"`
	AssetName            string `json:"asset_name"`
}
