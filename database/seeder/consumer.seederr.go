package seeder

import (
	"fmt"
	"gorm.io/gorm"
	"test-go-developer/database/entities"
	"time"
)

func SeedData(db *gorm.DB) error {
	var checkCustomers []entities.Customer
	result := db.Find(&checkCustomers)
	if result.Error != nil {
		panic(result.Error)
	}
	if len(checkCustomers) != 0 {
		return fmt.Errorf("Successfully Seeder")
	}
	// Create Customer seed data
	customers := []entities.Customer{
		{
			NIK:          "1234567890123456",
			FullName:     "Budi Setiawan",
			LegalName:    "Budi",
			DateOfBirth:  time.Date(1990, 1, 1, 0, 0, 0, 0, time.UTC),
			PlaceOfBirth: "Jakarta",
			Salary:       15000000,
			KtpPhoto:     "ktp_john_doe.jpg",
			SelfiePhoto:  "selfie_john_doe.jpg",
		},
		{
			NIK:          "6543210987654321",
			FullName:     "Annisa Fahma",
			LegalName:    "Annisa Fahma",
			DateOfBirth:  time.Date(1995, 5, 10, 0, 0, 0, 0, time.UTC),
			PlaceOfBirth: "Bandung",
			Salary:       10000000,
			KtpPhoto:     "ktp_jane_doe.jpg",
			SelfiePhoto:  "selfie_jane_doe.jpg",
		},
	}
	var createdIDs []string // Slice to hold created IDs
	for _, customer := range customers {
		if err := db.Create(&customer).Error; err != nil {
			return err
		}
		// Append the created customer's ID to the slice
		createdIDs = append(createdIDs, customer.ID)
	}

	// Create LoanLimit seed data
	loanLimits := []entities.LoanLimit{
		{
			Tenor:      1,
			Limit:      100000,
			CustomerID: createdIDs[0],
		}, {
			Tenor:      2,
			Limit:      200000,
			CustomerID: createdIDs[0],
		}, {

			Tenor:      3,
			Limit:      500000,
			CustomerID: createdIDs[0],
		}, {
			Tenor:      6,
			Limit:      7000000,
			CustomerID: createdIDs[0],
		},
		{
			Tenor:      1,
			Limit:      1000000,
			CustomerID: createdIDs[1],
		}, {

			Tenor:      2,
			Limit:      1200000,
			CustomerID: createdIDs[1],
		}, {
			Tenor:      3,
			Limit:      1500000,
			CustomerID: createdIDs[1],
		}, {
			Tenor:      6,
			Limit:      20000000,
			CustomerID: createdIDs[1],
		},
	}

	// Insert LoanLimit seed data
	for _, loanLimit := range loanLimits {
		if err := db.Create(&loanLimit).Error; err != nil {
			return err
		}
	}
	fmt.Println("successfully seeder done")
	return nil
}
