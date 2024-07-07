package services

import (
	"fmt"
	"sso-auth/app/helpers"
	"sso-auth/app/http/requests"
	"sso-auth/app/models"
	"sso-auth/app/repositories"
	"sso-auth/app/responses"
	"sso-auth/app/schemas"
	oauth2authorizeservices "sso-auth/app/services/oauth2_authorization_services"
	"sso-auth/app/utils"
	"sync"

	"github.com/google/uuid"
)

type OAuth2Service struct {
	authRepository            *repositories.AuthRepository
	passwordCredentialService *oauth2authorizeservices.PasswordCredetialService
	authCodeService           *oauth2authorizeservices.AuthCodeService
	txRepo                    *repositories.TxRepository
}

func NewOAuth2Service() *OAuth2Service {
	return &OAuth2Service{
		authRepository:            repositories.NewAuthRepository(),
		passwordCredentialService: oauth2authorizeservices.NewPasswordCredentialService(),
		authCodeService:           oauth2authorizeservices.NewAuthCodeService(),
		txRepo:                    repositories.NewTxRepository(),
	}
}

func (s *OAuth2Service) Login(request *requests.OAuth2LoginRequest) (res any, err error) {

	switch request.GrantType {
	case string(requests.GrantTypePasswordCredential):
		res, err = s.passwordCredentialService.Login(request)

	case string(requests.GrantTypeAuthCode):
		res, err = s.authCodeService.Login(request)
	}
	return
}

func (s *OAuth2Service) Register(data *requests.OAuth2Request) (*responses.OtpResponse, error) {

	err := s.authRepository.CheckEmailExists(data.Email, &models.User{})

	if err == nil {
		return nil, &schemas.ResponseApiError{
			Status:  schemas.ApiErrorBadRequest,
			Message: "Email already taken",
		}
	}

	user := models.User{
		Name:     data.Name,
		Password: data.Password,
		Email:    data.Email,
		Phone:    data.Phone,
	}

	err = s.authRepository.CreateUser(&user)

	if err != nil {
		return nil, &schemas.ResponseApiError{
			Status:  schemas.ApiErrorInternalServer,
			Message: "Internal server error",
		}
	}
	// generate otp
	otpHelper := helpers.NewOtp()
	otpCode := otpHelper.GenerateOTP()
	// send email
	uniqueCode := uuid.NewString()
	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		defer wg.Done()
		errMail := utils.NewEmailBuilder().To([]string{data.Email}).Content(utils.MailContent{
			Subject: "SSO Auth Email Verification",
			Html:    otpCode,
		}).Send()

		if errMail != nil {
			fmt.Println("[FAILED][SEND][OTP][EMAIL] %v \n", errMail)
		}

		_, errOtp := otpHelper.Save(user.ID, otpCode, uniqueCode)

		if errOtp != nil {
			fmt.Println("[FAILED][OTP][ERROR] %v \n", errOtp)
		}

	}()

	wg.Wait()
	return &responses.OtpResponse{
		UniqueCode: uniqueCode,
	}, nil
}

func (s *OAuth2Service) VerifOtp(request *requests.VerifOtp) (*responses.VerifOtpResponse, error) {
	var user models.User
	err := s.authRepository.CheckEmailExists(request.Email, &user)

	if err != nil {
		return nil, &schemas.ResponseApiError{
			Status:  schemas.ApiErrorBadRequest,
			Message: "Email not found",
		}
	}

	otpHelper := helpers.NewOtp()
	valid, err := otpHelper.IsOtpValid(request.UniqueCode, request.Otp)

	if !valid {
		return nil, &schemas.ResponseApiError{
			Status:  schemas.ApiErrorBadRequest,
			Message: err.Error(),
		}
	}

	tx := s.txRepo.Begin()
	defer tx.Commit()

	err = s.authRepository.UpdateWithTx(tx, &models.User{IsActive: true}, user.ID)

	if err != nil {
		s.txRepo.Rollback(tx)
		return nil, &schemas.ResponseApiError{
			Status:  schemas.ApiErrorInternalServer,
			Message: "Internal server error",
		}
	}

	return &responses.VerifOtpResponse{
		Name:  user.Name,
		Email: user.Email,
		Phone: user.Phone,
	}, nil
}
