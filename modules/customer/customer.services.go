package customer

import (
	"fmt"
	"gorm.io/gorm"
	"strings"
	"test-go-developer/database"
	"test-go-developer/database/entities"
	"test-go-developer/modules/customer/request"
	"test-go-developer/modules/customer/response"
	"time"
)

func createService(payload request.CustomerRequests) (interface{}, error) {
	var existingCustomer entities.Customer
	var customer entities.Customer

	// Start a new transaction
	tx := database.DB.Begin()

	// Check if a customer with the same NIK already exists
	if err := tx.Model(&entities.Customer{}).Where("nik = ?", payload.NIK).First(&existingCustomer).Error; err == nil {
		// Rollback transaction if customer already exists
		tx.Rollback()
		return nil, fmt.Errorf("customer with NIK %s already exists", payload.NIK)
	}
	dateOfBirth, _ := time.Parse("2006-01-02", payload.DateOfBirth)
	customer = entities.Customer{
		NIK:          payload.NIK,
		FullName:     payload.FullName,
		LegalName:    payload.LegalName,
		PlaceOfBirth: payload.PlaceOfBirth,
		DateOfBirth:  dateOfBirth, // Assuming this is already parsed
		Salary:       payload.Salary,
		KtpPhoto:     payload.KtpPhoto,
		SelfiePhoto:  payload.SelfiePhoto,
	}

	salary := payload.Salary // Assuming salary is of type int64
	limitPerMonth := salary * 30 / 100
	// Create the customer
	if err := tx.Create(&customer).Error; err != nil {
		tx.Rollback() // Rollback transaction on error
		return nil, err
	}
	// Create loan limits in a loop
	for i := 1; i <= 4; i++ {
		if err := createLoanLimit(tx, customer.ID, limitPerMonth*int64(i), i); err != nil {
			tx.Rollback() // Rollback transaction on error
			return nil, err
		}
	}

	// Commit the transaction
	if err := tx.Commit().Error; err != nil {
		return nil, err // Return error if commit fails
	}

	return customer, nil
}
func createLoanLimit(tx *gorm.DB, customerID string, limit int64, tenor int) error {
	loanLimit := entities.LoanLimit{
		CustomerID: customerID,
		Limit:      limit,
		Tenor:      tenor,
	}
	return tx.Create(&loanLimit).Error // Use the transaction object here
}

func getCustomerService(query request.CustomerQuery) (interface{}, error) {
	var customers []entities.Customer

	// Build the query
	dbQuery := database.DB.Unscoped().Preload("LoanLimit").Order("created_at DESC")

	// Apply filters based on the query parameters
	if query.FullName != "" {
		dbQuery = dbQuery.Where("LOWER(full_name) LIKE ?", strings.ToLower(query.FullName))
	}
	if query.NIK != "" {
		dbQuery = dbQuery.Where("nik LIKE ?", query.NIK)
	}

	// Execute the query and check for errors
	if err := dbQuery.Find(&customers).Error; err != nil {
		return nil, err // Return the error if the query fails
	}

	// Map the customer entities to response format
	var customerResponses []response.CustomersResponse
	for _, customer := range customers {
		var loanLimits []response.LoanLimit
		for _, limit := range customer.LoanLimit {
			loanLimits = append(loanLimits, response.LoanLimit{
				ID:    limit.ID,
				Tenor: limit.Tenor,
				Limit: limit.Limit,
			})
		}

		customerResponses = append(customerResponses, response.CustomersResponse{
			ID:           customer.ID,
			NIK:          customer.NIK,
			FullName:     customer.FullName,
			LegalName:    customer.LegalName,
			DateOfBirth:  customer.DateOfBirth,
			PlaceOfBirth: customer.PlaceOfBirth,
			Salary:       customer.Salary,
			KtpPhoto:     customer.KtpPhoto,
			SelfiePhoto:  customer.SelfiePhoto,
			LoanLimit:    loanLimits,
		})
	}

	return customerResponses, nil // Return the list of customer responses
}
