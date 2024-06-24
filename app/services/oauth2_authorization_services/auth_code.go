package oauth2authorizationservices

import (
	"sso-auth/app/facades"
	"sso-auth/app/http/requests"
	"sso-auth/app/models"
	"sso-auth/app/repositories"
	"sso-auth/app/responses"
	"sso-auth/app/schemas"
	"sso-auth/app/utils"
	"time"
)

type AuthCodeService struct {
	clientRepository   *repositories.ClientRepository
	authRepository     *repositories.AuthRepository
	authCodeRepository *repositories.AuthCodeRepository
}

func NewAuthCodeService() *AuthCodeService {
	return &AuthCodeService{
		clientRepository:   repositories.NewClientRepository(),
		authRepository:     repositories.NewAuthRepository(),
		authCodeRepository: repositories.NewAuthCodeRepository(),
	}
}

func (s *AuthCodeService) Login(request requests.OAuth2LoginRequest) (*responses.LoginResponsesAuthCode, error) {
	var clientData models.Client
	err := s.clientRepository.GetById(&clientData, request.ClientId)

	if err != nil {
		return nil, &schemas.ResponseApiError{
			Status:  schemas.ApiErrorNotFound,
			Message: "Client not found",
		}
	}

	var userData models.User
	err = s.authRepository.GetUser(&userData, request.Username)
	if err != nil {
		return nil, &schemas.ResponseApiError{
			Status:  schemas.ApiErrorForbidden,
			Message: "Username or password invalid",
		}
	}

	auth := utils.CheckPasswordHash(request.Password, userData.Password)

	if !auth {
		return nil, &schemas.ResponseApiError{
			Status:  schemas.ApiErrorForbidden,
			Message: "Username or password invalid",
		}
	}

	authCode := facades.RandomString(64)

	authCodeData := models.AuthCode{
		Code:       &authCode,
		ExpiryTime: time.Now().Add(time.Minute * 10),
		ClientId:   clientData.ID,
		UserId:     userData.ID,
	}
	err = s.authCodeRepository.Create(&authCodeData)

	res := responses.LoginResponsesAuthCode{
		Id:       authCodeData.UserId,
		Username: userData.Username,
		Email:    userData.Email,
		AuthCode: responses.AuthCode{
			Code:       *authCodeData.Code,
			ExpiryTime: authCodeData.ExpiryTime,
		},
	}

	if err != nil {
		return nil, &schemas.ResponseApiError{
			Status:  schemas.ApiErrorForbidden,
			Message: err.Error(),
		}
	}

	return &res, nil
}
