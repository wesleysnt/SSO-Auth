package services

import (
	"sso-auth/app/http/requests"
	"sso-auth/app/models"
	"sso-auth/app/repositories"
	"sso-auth/app/responses"
	oauth2authorizationservices "sso-auth/app/services/oauth2_authorization_services"
)

type OAuth2Service struct {
	authRepository            *repositories.AuthRepository
	passwordCredentialService *oauth2authorizationservices.PasswordCredetialService
	authCodeService           *oauth2authorizationservices.AuthCodeService
}

func NewOAuth2Service() *OAuth2Service {
	return &OAuth2Service{
		authRepository:            repositories.NewAuthRepository(),
		passwordCredentialService: oauth2authorizationservices.NewPasswordCredentialService(),
	}
}

func (s *OAuth2Service) Login(grantType, redirectUri string, request *requests.OAuth2LoginRequest) (res *responses.LoginResponses, err error) {

	switch grantType {
	case string(requests.GrantTypePasswordCredential):
		res, err = s.passwordCredentialService.Login(*request)

	case string(requests.GrantTypeAuthCode):
		break
	}
	return
}

func (s *OAuth2Service) Register(data *requests.OAuth2Request) (err error) {

	user := models.User{
		Username: data.Username,
		Password: data.Password,
		Email:    data.Email,
	}

	err = s.authRepository.CreateUser(&user)

	return
}
