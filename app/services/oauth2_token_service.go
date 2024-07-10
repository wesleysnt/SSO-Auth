package services

import (
	"sso-auth/app/facades"
	"sso-auth/app/http/requests"
	"sso-auth/app/models"
	"sso-auth/app/repositories"
	"sso-auth/app/responses"
	"sso-auth/app/schemas"
	oauth2authorizationservices "sso-auth/app/services/oauth2_authorization_services"
)

type Oauth2TokenService struct {
	clientRepository       *repositories.ClientRepository
	authRepository         *repositories.AuthRepository
	authCodeService        *oauth2authorizationservices.AuthCodeService
	clientCredential       *oauth2authorizationservices.ClientCredentialService
	refreshTokenRepository *repositories.RefreshTokenRepository
	accessTokenRepository  *repositories.AccessTokenRepository
	passwordCredential     *oauth2authorizationservices.PasswordCredetialService
}

func NewOauth2TokenService() *Oauth2TokenService {
	return &Oauth2TokenService{
		clientRepository:       repositories.NewClientRepository(),
		authRepository:         repositories.NewAuthRepository(),
		authCodeService:        oauth2authorizationservices.NewAuthCodeService(),
		clientCredential:       oauth2authorizationservices.NewClientCredentialService(),
		refreshTokenRepository: repositories.NewRefreshTokenRepository(),
		accessTokenRepository:  repositories.NewAccessTokenRepository(),
		passwordCredential:     oauth2authorizationservices.NewPasswordCredentialService(),
	}
}

func (s *Oauth2TokenService) Token(request *requests.TokenRequest) (res *responses.TokenResponse, err error) {
	switch request.GrantType {
	case requests.GrantTypeAuthCode:
		res, err = s.authCodeService.Token(request)
	case requests.GrantTypeClientCredential:
		res, err = s.clientCredential.Token(request)
	}

	return

}

func (s *Oauth2TokenService) ValidateToken(request *requests.ValidateTokenRequest) (res *responses.ValidateTokenResponse, err error) {

	switch request.GrantType {
	case requests.GrantTypeAuthCode:
		res, err = s.authCodeService.ValidateToken(request)
	case requests.GrantTypeClientCredential:
		res, err = s.clientCredential.ValidateToken(request)
	case requests.GrantTypePasswordCredential:
		res, err = s.passwordCredential.ValidateToken(request)
	}

	return
}

func (s *Oauth2TokenService) RefreshToken(request *requests.RefreshTokenRequest) (*responses.TokenResponse, error) {
	var clientData models.Client
	err := s.clientRepository.GetByClientId(&clientData, request.ClientId)

	if err != nil {
		return nil, &schemas.ResponseApiError{
			Status:  schemas.ApiErrorNotFound,
			Message: "Client not found!",
		}
	}

	var userData models.User

	err = s.authRepository.GetById(&userData, request.UserId)
	if err != nil {
		return nil, &schemas.ResponseApiError{
			Status:  schemas.ApiErrorNotFound,
			Message: "User not found!",
		}
	}
	err = s.refreshTokenRepository.Check(request.RefreshToken, clientData.ID, request.UserId)

	if err != nil {
		return nil, &schemas.ResponseApiError{
			Status:  schemas.ApiErrorNotFound,
			Message: "Token not invalid or not found",
		}
	}

	token, err := facades.ParseToken(request.RefreshToken, request.Secret)

	if err != nil {
		return nil, &schemas.ResponseApiError{
			Status:  schemas.ApiErrorUnprocessAble,
			Message: err.Error(),
		}
	}

	clientId, _ := facades.GetClientIdFromToken(request.RefreshToken, request.Secret)
	userId, _ := facades.GetUserIdFromToken(request.RefreshToken, request.Secret)

	// Generate access token
	tokenString, err := facades.GenerateToken(request.Secret, clientId, userId, 2)

	if err != nil {
		return nil, &schemas.ResponseApiError{
			Status:  schemas.ApiErrorUnprocessAble,
			Message: "Something went wrong when generating access token",
		}
	}

	token, err = facades.ParseToken(tokenString, request.Secret)

	if err != nil {
		return nil, &schemas.ResponseApiError{
			Status:  schemas.ApiErrorUnprocessAble,
			Message: err.Error(),
		}
	}

	tokenExpired, _ := token.Claims.GetExpirationTime()

	errSaveToken := s.accessTokenRepository.Create(&models.AccessToken{
		Token:      tokenString,
		UserId:     userId,
		ClientId:   clientData.ID,
		ExpiryTime: tokenExpired.Time,
	})

	if errSaveToken != nil {
		return nil, &schemas.ResponseApiError{
			Status:  schemas.ApiErrorUnprocessAble,
			Message: errSaveToken.Error(),
		}
	}

	// Generate Refresh Token

	refreshTokenString, err := facades.GenerateToken(request.Secret, clientId, userId, 4)

	if err != nil {
		return nil, &schemas.ResponseApiError{
			Status:  schemas.ApiErrorUnprocessAble,
			Message: "Something went wrong when generating refresh token",
		}
	}

	refreshTokenExpired, _ := token.Claims.GetExpirationTime()

	errSaveRefreshToken := s.refreshTokenRepository.Create(&models.RefreshToken{
		Token:      refreshTokenString,
		UserId:     userId,
		ClientId:   clientData.ID,
		ExpiryTime: refreshTokenExpired.Time,
	})

	if errSaveRefreshToken != nil {
		return nil, &schemas.ResponseApiError{
			Status:  schemas.ApiErrorUnprocessAble,
			Message: errSaveRefreshToken.Error(),
		}
	}

	res := responses.TokenResponse{

		AccessToken: responses.AccessToken{
			Token:      tokenString,
			ExpiryTime: tokenExpired,
		},
		RefreshToken: responses.RefreshToken{
			Token:      refreshTokenString,
			ExpiryTime: refreshTokenExpired,
		},
	}

	return &res, err
}
