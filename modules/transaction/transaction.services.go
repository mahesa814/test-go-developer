package transaction

import (
	"fmt"
	"math/rand"
	"strings"
	"test-go-developer/database"
	"test-go-developer/database/entities"
	responseCustomer "test-go-developer/modules/customer/response"
	"test-go-developer/modules/transaction/request"
	"test-go-developer/modules/transaction/response"
	"time"
)

const interestRate = 0.05 // 5%

func createTransactionService(payload request.TransactionRequest) (interface{}, error) {
	// Start a new transaction
	tx := database.DB.Begin()

	var customer entities.Customer
	if err := tx.Where("id = ? ", payload.CustomerID).First(&customer).Error; err != nil {
		tx.Rollback() // Rollback transaction if customer not found
		return nil, fmt.Errorf("customer not found")
	}

	var loanLimit entities.LoanLimit
	if err := tx.Where("customer_id = ? AND tenor = ?", payload.CustomerID, payload.Tenor).First(&loanLimit).Error; err != nil {
		tx.Rollback() // Rollback transaction if loan limit not found
		return nil, fmt.Errorf("loan limit not found")
	}

	// Validate if asset price exceeds the limit
	if float64(payload.AssetPrice) > float64(loanLimit.Limit) {
		// Create a detailed error message
		return nil, fmt.Errorf("maaf, limit anda tidak cukup untuk transaksi ini. Tenor: %d bulan, sisa limit: %d", payload.Tenor, loanLimit.Limit)
	}

	// Calculate OTR, admin fee, installment amount, and interest amount
	otr := float64(payload.AssetPrice)   // OTR taken from asset price
	adminFee := 0.03 * otr               // Example admin fee 3%
	interestAmount := otr * interestRate // Calculate interest
	installment := (otr + adminFee + interestAmount) / float64(payload.Tenor)

	transaction := entities.Transaction{
		ContractNumber: generateContractNumber(), // Function to generate contract number
		OTR:            otr,
		AdminFee:       adminFee,
		Installment:    installment,
		InterestAmount: interestAmount,
		AssetName:      payload.AssetName, // Asset name being purchased
		LoanLimitId:    loanLimit.ID,
		CustomerID:     customer.ID,
	}

	// Save transaction to the database
	if err := tx.Preload("Customer").Create(&transaction).Error; err != nil {
		tx.Rollback() // Rollback transaction on error
		return nil, err
	}
	loanLimit.Limit -= payload.AssetPrice
	if err := tx.Save(&loanLimit).Error; err != nil {
		tx.Rollback() // Rollback transaction on error
		return nil, err
	}
	// Commit the transaction
	if err := tx.Commit().Error; err != nil {
		return nil, err // Return error if commit fails
	}
	transactionResponse := response.TransactionsResponse{ID: transaction.ID,
		ContractNumber: transaction.ContractNumber,
		OTR:            transaction.OTR,
		AdminFee:       transaction.AdminFee,
		InterestAmount: transaction.InterestAmount,
		AssetName:      transaction.AssetName,
		Installment:    transaction.Installment,
		CreatedAt:      transaction.CreatedAt,
	}
	return transactionResponse, nil
}

func generateContractNumber() string {
	// Set seed untuk random number generator
	rand.Seed(time.Now().UnixNano())

	// Generate random number, misalnya antara 10000 dan 99999
	randomNumber := rand.Intn(90000) + 10000

	// Gabungkan nama depan dengan nomor acak
	contractNumber := fmt.Sprintf("MAHESA%d", randomNumber)

	return contractNumber
}

func getTransactionsService(query request.TransactionQuery) (interface{}, error) {
	var transactions []entities.Transaction // Renamed variable to 'customers' for clarity

	// Build the query
	dbQuery := database.DB.Unscoped().Preload("Customer").Preload("LoanLimit").Order("created_at DESC")

	// Apply filters based on the query parameters
	if query.ContractNumber != "" {
		dbQuery = dbQuery.Where("LOWER(contract_number) LIKE ?", strings.ToLower(query.ContractNumber))
	}
	if query.AssetName != "" {
		dbQuery = dbQuery.Where("LOWER(asset_name) LIKE ?", strings.ToLower(query.AssetName))
	}

	// Execute the query and check for errors
	if err := dbQuery.Find(&transactions).Error; err != nil {
		return nil, err // Return the error if the query fails
	}

	var transactionResponses []response.TransactionsResponse
	for _, transaction := range transactions {
		transactionResponses = append(transactionResponses, response.TransactionsResponse{
			ID:             transaction.ID,
			ContractNumber: transaction.ContractNumber,
			OTR:            transaction.OTR,
			AdminFee:       transaction.AdminFee,
			Installment:    transaction.Installment,
			InterestAmount: transaction.InterestAmount,
			AssetName:      transaction.AssetName,
			Customer: &responseCustomer.CustomersResponse{
				ID:           transaction.CustomerID,
				FullName:     transaction.Customer.FullName,
				NIK:          transaction.Customer.NIK,
				LegalName:    transaction.Customer.LegalName,
				KtpPhoto:     transaction.Customer.KtpPhoto,
				SelfiePhoto:  transaction.Customer.SelfiePhoto,
				PlaceOfBirth: transaction.Customer.PlaceOfBirth,
				Salary:       transaction.Customer.Salary,
			},
		})
	}

	return transactionResponses, nil // Return the list of customers
}
