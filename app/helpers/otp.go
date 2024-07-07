package helpers

import (
	"crypto/rand"
	"fmt"
	"math"
	"math/big"
	"os"
	"sso-auth/app/models"
	"sso-auth/app/repositories"
	"strconv"
	"time"

	"github.com/goravel/framework/facades"
)

type Otp struct {
	otpRequestRepo *repositories.OtpRequestRepository
}

func NewOtp() Otp {
	return Otp{otpRequestRepo: repositories.NewOtpRequestRepository()}
}

func (h *Otp) Save(userId uint, otpCode, UniqueCode string) (*models.OtpRequests, error) {
	var data models.OtpRequests
	periode := os.Getenv("OTP_PERIODE")
	convPeriode, _ := strconv.Atoi(periode)

	data.OtpCode = otpCode
	data.IsUsed = false
	data.UniqueCode = UniqueCode
	data.ExpiredAt = time.Now().Add(time.Second * time.Duration(convPeriode+60))
	data.UserId = userId

	return h.otpRequestRepo.Store(&data)
}

func (h *Otp) Resend(userId uint, otpCode, UniqueCode string) (*models.OtpRequests, error) {
	var data models.OtpRequests
	periode := os.Getenv("OTP_PERIODE")
	convPeriode, _ := strconv.Atoi(periode)

	data.OtpCode = otpCode
	data.IsUsed = false
	data.UniqueCode = UniqueCode
	data.ExpiredAt = time.Now().Add(time.Second * time.Duration(convPeriode+60))
	data.UserId = userId

	// delete data
	err := h.deleteByCustId(userId)
	if err != nil {
		facades.Log().Infof("[Resend][OTP][DELETE][ERROR] %v", err)
	}

	return h.otpRequestRepo.Store(&data)
}

func (h *Otp) delete(uniqueCode string) error {
	err := h.otpRequestRepo.Delete(uniqueCode)
	return err
}

func (h *Otp) deleteByCustId(custId uint) error {
	err := h.otpRequestRepo.DeleteByCustId(custId)
	return err
}

func (h *Otp) GenerateOTP() string {
	digit := os.Getenv("OTP_DIGIT")
	convDigit, _ := strconv.Atoi(digit)
	maxDigits := uint32(convDigit)

	bi, _ := rand.Int(
		rand.Reader,
		big.NewInt(int64(math.Pow(10, float64(maxDigits)))),
	)

	return fmt.Sprintf("%0*d", maxDigits, bi)
}

func (h *Otp) IsOtpValid(uniqueCode, otp string) (bool, error) {
	data, errData := h.otpRequestRepo.GetByUniqueCode(uniqueCode)

	if errData != nil {
		return false, fmt.Errorf("credentials invalid")
	}

	now := time.Now()

	// expired token
	if now.After(data.ExpiredAt) {
		return false, fmt.Errorf("otp expired")
	}

	if data.OtpCode != otp {
		return false, fmt.Errorf("invalid otp")
	}

	// delete data
	errDelete := h.delete(uniqueCode)

	if errDelete != nil {
		facades.Log().Infof("[IsOtpValid][DELETE][ERROR] %v", errDelete)
	}

	return true, nil
}
