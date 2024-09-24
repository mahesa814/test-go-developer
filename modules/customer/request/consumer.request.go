package request

type CustomerQuery struct {
	NIK      string `form:"nik"`
	FullName string `form:"full_name"`
}

type CustomerRequests struct {
	NIK          string `json:"nik" `
	FullName     string `json:"full_name" `
	LegalName    string `json:"legal_name"`
	PlaceOfBirth string `json:"place_of_birth" `
	DateOfBirth  string `json:"date_of_birth"`
	Salary       int64  `json:"salary"`
	KtpPhoto     string `json:"ktp_photo" `
	SelfiePhoto  string `json:"selfie_photo" `
}
