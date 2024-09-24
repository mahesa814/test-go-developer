package request

type TransactionQuery struct {
	ContractNumber string `form:"contract_number" json:"contract_number"`
	AssetName      string `form:"asset_name" json:"asset_name"`
}

type TransactionRequest struct {
	AssetName  string `json:"asset_name" binding:"required" required:"$field is required"`
	AssetPrice int64  `json:"asset_price" binding:"required" required:"$field is required"`
	Tenor      int    `json:"tenor" binding:"required" required:"$field is required"`
	CustomerID string `json:"customer_id" binding:"required" required:"$field is required"`
}
