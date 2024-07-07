package responses

type OtpResponse struct {
	UniqueCode string `json:"unique_code"`
}

type VerifOtpResponse struct {
	Name  string `json:"name"`
	Email string `json:"email"`
	Phone string `json:"phone"`
}
