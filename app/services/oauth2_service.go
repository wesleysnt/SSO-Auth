package services

import (
	"context"
	"fmt"
	"sso-auth/app/facades"
	"sso-auth/app/helpers"
	"sso-auth/app/http/requests"
	"sso-auth/app/models"
	"sso-auth/app/repositories"
	"sso-auth/app/responses"
	"sso-auth/app/schemas"
	oauth2authorizeservices "sso-auth/app/services/oauth2_authorization_services"
	"sso-auth/app/utils"
	"sync"
	"time"

	"github.com/google/uuid"
)

type OAuth2Service struct {
	authRepository            *repositories.AuthRepository
	passwordCredentialService *oauth2authorizeservices.PasswordCredetialService
	authCodeService           *oauth2authorizeservices.AuthCodeService
	txRepo                    *repositories.TxRepository
	accessTokenRepository     *repositories.AccessTokenRepository
	clientRepository          *repositories.ClientRepository
	authCodeRepository        *repositories.AuthCodeRepository
	forgotPasswordToken       *repositories.ForgotPasswordTokenRepository
}

func NewOAuth2Service() *OAuth2Service {
	return &OAuth2Service{
		authRepository:            repositories.NewAuthRepository(),
		passwordCredentialService: oauth2authorizeservices.NewPasswordCredentialService(),
		authCodeService:           oauth2authorizeservices.NewAuthCodeService(),
		txRepo:                    repositories.NewTxRepository(),
		accessTokenRepository:     repositories.NewAccessTokenRepository(),
		clientRepository:          repositories.NewClientRepository(),
		authCodeRepository:        repositories.NewAuthCodeRepository(),
		forgotPasswordToken:       repositories.NewForgotPasswordTokenRepository(),
	}
}

func (s *OAuth2Service) Login(request *requests.OAuth2LoginRequest, ctx context.Context) (res any, err error) {

	switch request.GrantType {
	case requests.GrantTypePasswordCredential:
		res, err = s.passwordCredentialService.Login(request, ctx)

	case requests.GrantTypeAuthCode:
		res, err = s.authCodeService.Login(request, ctx)
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

	isPhoneValid := utils.IsPhoneNumber(data.Phone)

	if !isPhoneValid {
		return nil, &schemas.ResponseApiError{
			Status:  schemas.ApiErrorBadRequest,
			Message: "Phone number is invalid",
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
			Html:    fmt.Sprintf("Use this otp code to verify your email \n OTP: %v", otpCode),
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

func (s *OAuth2Service) IsLoggedIn(request *requests.IsLoggedInRequest, ctx context.Context) (*responses.IsLoggedInResponse, error) {
	var clientData models.Client

	err := s.clientRepository.GetByClientId(&clientData, request.ClientId, ctx)

	if err != nil {
		return nil, &schemas.ResponseApiError{
			Status:  schemas.ApiErrorBadRequest,
			Message: "Client not found!",
		}
	}

	accessToken, err := s.accessTokenRepository.GetByTokenAndClient(request.Token, clientData.ID)

	if err != nil {
		return nil, &schemas.ResponseApiError{
			Status:  schemas.ApiErrorBadRequest,
			Message: "Token not found!",
		}
	}

	_, err = facades.ParseToken(accessToken.Token, clientData.Secret)
	if err != nil {
		return nil, &schemas.ResponseApiError{
			Status:  schemas.ApiErrorBadRequest,
			Message: err.Error(),
		}
	}

	userId, _ := facades.GetUserIdFromToken(accessToken.Token, clientData.Secret)

	authCode := facades.RandomString(64)

	authCodeData := models.AuthCode{
		Code:       &authCode,
		ExpiryTime: time.Now().Add(time.Minute * 10),
		ClientId:   clientData.ID,
		UserId:     userId,
	}
	err = s.authCodeRepository.Create(&authCodeData)

	if err != nil {
		return nil, &schemas.ResponseApiError{
			Status:  schemas.ApiErrorUnprocessAble,
			Message: err.Error(),
		}
	}

	return &responses.IsLoggedInResponse{
		IsLoggedIn: true,
		AuthCode:   authCode,
	}, nil
}

func (s *OAuth2Service) RequestForgotPassword(request *requests.RequestForgotPasswordRequest) error {
	var userData models.User
	err := s.authRepository.CheckEmailExists(request.Email, &userData)

	if err != nil {
		return &schemas.ResponseApiError{
			Status:  schemas.ApiErrorBadRequest,
			Message: "User not found!",
		}
	}

	token := uuid.NewString()

	err = s.forgotPasswordToken.Create(&models.ForgotPasswordToken{
		UserId:     userData.ID,
		Token:      token,
		ExpiryTime: time.Now().Add(time.Minute * 10),
	})

	if err != nil {
		return &schemas.ResponseApiError{
			Status:  schemas.ApiErrorBadRequest,
			Message: err.Error(),
		}
	}

	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		defer wg.Done()
		errMail := utils.NewEmailBuilder().To([]string{request.Email}).Content(utils.MailContent{
			Subject: "SSO Auth Email Verification",
			Html:    fmt.Sprintf("Open the link below to reset your password. \n Link/%v", token),
		}).Send()

		if errMail != nil {
			fmt.Println("[FAILED][SEND][OTP][EMAIL] %v \n", errMail)
		}
	}()

	wg.Wait()
	return nil
}

func (s *OAuth2Service) ResetPassword(request *requests.ResetPasswordRequest) error {
	var forgotPasswordToken models.ForgotPasswordToken
	err := s.forgotPasswordToken.FindByToken(request.Token, &forgotPasswordToken)

	if err != nil {
		return &schemas.ResponseApiError{
			Status:  schemas.ApiErrorBadRequest,
			Message: "Token not found!",
		}
	}

	if forgotPasswordToken.ExpiryTime.Before(time.Now()) {
		return &schemas.ResponseApiError{
			Status:  schemas.ApiErrorBadRequest,
			Message: "Token expired!",
		}
	}

	err = s.authRepository.UpdatePassword(forgotPasswordToken.UserId, request.Password)

	if err != nil {
		return &schemas.ResponseApiError{
			Status:  schemas.ApiErrorBadRequest,
			Message: err.Error(),
		}
	}

	err = s.forgotPasswordToken.UpdateIsUsed(forgotPasswordToken.Token)
	if err != nil {
		return &schemas.ResponseApiError{
			Status:  schemas.ApiErrorBadRequest,
			Message: err.Error(),
		}
	}
	return nil
}
