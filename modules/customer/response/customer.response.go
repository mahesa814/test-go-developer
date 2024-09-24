package response

import "time"

type CustomersResponse struct {
	ID           string      `json:"id"`
	NIK          string      `json:"nik"`
	FullName     string      `json:"full_name"`
	LegalName    string      `json:"legal_name"`
	DateOfBirth  time.Time   `json:"date_of_birth"`
	PlaceOfBirth string      `json:"place_of_birth"`
	Salary       int64       `json:"salary"`
	KtpPhoto     string      `json:"ktp_photo"`
	SelfiePhoto  string      `json:"selfie_photo"`
	LoanLimit    []LoanLimit `json:"loan_limits"`
}
type LoanLimit struct {
	ID    string `json:"id"`
	Tenor int    `json:"tenor"`
	Limit int64  `limit:"limit"`
}
