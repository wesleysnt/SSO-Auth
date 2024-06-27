package services

import (
	"sso-auth/app/http/requests"
	"sso-auth/app/models"
	"sso-auth/app/repositories"
	oauth2authorizeservices "sso-auth/app/services/oauth2_authorization_services"
)

type OAuth2Service struct {
	authRepository            *repositories.AuthRepository
	passwordCredentialService *oauth2authorizeservices.PasswordCredetialService
	authCodeService           *oauth2authorizeservices.AuthCodeService
}

func NewOAuth2Service() *OAuth2Service {
	return &OAuth2Service{
		authRepository:            repositories.NewAuthRepository(),
		passwordCredentialService: oauth2authorizeservices.NewPasswordCredentialService(),
		authCodeService:           oauth2authorizeservices.NewAuthCodeService(),
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

func (s *OAuth2Service) Register(data *requests.OAuth2Request) (err error) {

	user := models.User{
		Username: data.Username,
		Password: data.Password,
		Email:    data.Email,
	}

	err = s.authRepository.CreateUser(&user)

	return
}
