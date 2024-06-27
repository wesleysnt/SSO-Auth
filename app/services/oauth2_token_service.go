package services

import (
<<<<<<< HEAD
	"sso-auth/app/facades"
	"sso-auth/app/http/requests"
	"sso-auth/app/models"
	"sso-auth/app/repositories"
	"sso-auth/app/responses"
	"sso-auth/app/schemas"
=======
	"sso-auth/app/http/requests"
	"sso-auth/app/responses"
<<<<<<< HEAD
>>>>>>> 325f9fc (.)
=======
>>>>>>> ec32a2f (.)
>>>>>>> 7a7a01f (.)
	oauth2authorizationservices "sso-auth/app/services/oauth2_authorization_services"
)

type Oauth2TokenService struct {
<<<<<<< HEAD
	authCodeService        *oauth2authorizationservices.AuthCodeService
	clientCredential       *oauth2authorizationservices.ClientCredentialService
	refreshTokenRepository *repositories.RefreshTokenRepository
	accessTokenRepository  *repositories.AccessTokenRepository
=======
	authCodeService  *oauth2authorizationservices.AuthCodeService
	clientCredential *oauth2authorizationservices.ClientCredentialService
<<<<<<< HEAD
>>>>>>> 325f9fc (.)
=======
>>>>>>> ec32a2f (.)
>>>>>>> 7a7a01f (.)
}

func NewOauth2TokenService() *Oauth2TokenService {
	return &Oauth2TokenService{
<<<<<<< HEAD
		authCodeService:        oauth2authorizationservices.NewAuthCodeService(),
		clientCredential:       oauth2authorizationservices.NewClientCredentialService(),
		refreshTokenRepository: repositories.NewRefreshTokenRepository(),
		accessTokenRepository:  repositories.NewAccessTokenRepository(),
	}
}

func (s *Oauth2TokenService) Token(request *requests.TokenRequest) (res *responses.TokenResponse, err error) {
	switch request.GrantType {
=======
		authCodeService:  oauth2authorizationservices.NewAuthCodeService(),
		clientCredential: oauth2authorizationservices.NewClientCredentialService(),
	}
}

func (s *Oauth2TokenService) Token(request *requests.TokenRequest, grantType, redirectUri string) (res *responses.TokenResponse, err error) {
	switch grantType {
<<<<<<< HEAD
>>>>>>> 325f9fc (.)
=======
>>>>>>> ec32a2f (.)
>>>>>>> 7a7a01f (.)
	case string(requests.GrantTypeAuthCode):
		res, err = s.authCodeService.Token(request)
	case string(requests.GrantTypeClientCredential):
		res, err = s.clientCredential.Token(request)
	}

	return

}
<<<<<<< HEAD

func (s *Oauth2TokenService) ValidateToken(request *requests.ValidateTokenRequest) (res responses.ValidateTokenResponse, err error) {
	token, err := facades.ParseToken(request.Token, request.Secret)

	if err != nil {
		return
	}

	tokenExp, _ := token.Claims.GetExpirationTime()
	clientId, _ := facades.GetClientIdFromToken(request.Token, request.Secret)
	userId, _ := facades.GetUserIdFromToken(request.Token, request.Secret)
	res = responses.ValidateTokenResponse{
		Active:   true,
		Exp:      *tokenExp,
		ClientId: clientId,
		UserId:   userId,
	}

	return
}

func (s *Oauth2TokenService) RefreshToken(request *requests.RefreshTokenRequest) (*responses.TokenResponse, error) {
	err := s.refreshTokenRepository.Check(request.RefreshToken, request.ClientId, request.UserId)

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
	tokenString, err := facades.GenerateToken(request.Secret, userId, clientId, 2)

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
		ClientId:   clientId,
		ExpiryTime: tokenExpired.Time,
	})

	if errSaveToken != nil {
		return nil, &schemas.ResponseApiError{
			Status:  schemas.ApiErrorUnprocessAble,
			Message: errSaveToken.Error(),
		}
	}

	// Generate Refresh Token

	refreshTokenString, err := facades.GenerateToken(request.Secret, userId, clientId, 4)

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
		ClientId:   clientId,
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
=======
<<<<<<< HEAD
>>>>>>> 325f9fc (.)
=======
>>>>>>> ec32a2f (.)
>>>>>>> 7a7a01f (.)
