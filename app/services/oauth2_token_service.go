package services

import (
	"sso-auth/app/facades"
	"sso-auth/app/http/requests"
	"sso-auth/app/responses"
	oauth2authorizationservices "sso-auth/app/services/oauth2_authorization_services"

	"github.com/golang-jwt/jwt/v5"
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

func (s *Oauth2TokenService) Token(request *requests.TokenRequest) (res *responses.TokenResponse, err error) {
	switch request.GrantType {
	case string(requests.GrantTypeAuthCode):
		res, err = s.authCodeService.Token(request)
	case string(requests.GrantTypeClientCredential):
		res, err = s.clientCredential.Token(request)
	}

	return

}

func (s *Oauth2TokenService) ValidateToken(request *requests.ValidateTokenRequest) (token *jwt.Token, err error) {
	token, err = facades.ParseToken(request.Token, request.Secret)

	if err != nil {
		return
	}

	return
}
