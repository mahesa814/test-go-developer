package migration

import (
	"test-go-developer/database"
	"test-go-developer/database/entities"
)

func AutoMigrate() error {
	// Declare the error variable
	var err error

	// Perform the migration and check for errors
	err = database.DB.AutoMigrate(
		&entities.Customer{},
		&entities.Transaction{},
		&entities.LoanLimit{},
	)
	if err != nil {
		return err // Return the error if it occurs
	}

	return nil // Return nil if no error occurs
}
