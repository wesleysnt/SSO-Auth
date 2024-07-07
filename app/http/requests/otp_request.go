package requests

type VerifOtp struct {
	Email      string `form:"email" json:"email"`
	Otp        string `form:"otp" json:"otp"`
	UniqueCode string `form:"unique_code" json:"unique_code"`
}
