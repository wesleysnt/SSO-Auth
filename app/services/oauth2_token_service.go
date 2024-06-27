package services

import (
	"sso-auth/app/http/requests"
	"sso-auth/app/responses"
	oauth2authorizationservices "sso-auth/app/services/oauth2_authorization_services"
)

type Oauth2TokenService struct {
	authCodeService  *oauth2authorizationservices.AuthCodeService
	clientCredential *oauth2authorizationservices.ClientCredentialService
}

func NewOauth2TokenService() *Oauth2TokenService {
	return &Oauth2TokenService{
		authCodeService:  oauth2authorizationservices.NewAuthCodeService(),
		clientCredential: oauth2authorizationservices.NewClientCredentialService(),
	}
}

func (s *Oauth2TokenService) Token(request *requests.TokenRequest, grantType, redirectUri string) (res *responses.TokenResponse, err error) {
	switch grantType {
	case string(requests.GrantTypeAuthCode):
		res, err = s.authCodeService.Token(request)
	case string(requests.GrantTypeClientCredential):
		res, err = s.clientCredential.Token(request)
	}

	return

}
