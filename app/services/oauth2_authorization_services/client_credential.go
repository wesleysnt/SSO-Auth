package oauth2authorizationservices

import (
	"context"
	"sso-auth/app/facades"
	"sso-auth/app/http/requests"
	"sso-auth/app/models"
	"sso-auth/app/repositories"
	"sso-auth/app/responses"
	"sso-auth/app/schemas"
)

type ClientCredentialService struct {
	accessTokenRepository  *repositories.AccessTokenRepository
	clientRepository       *repositories.ClientRepository
	refreshTokenRepository *repositories.RefreshTokenRepository
}

func NewClientCredentialService() *ClientCredentialService {
	return &ClientCredentialService{
		accessTokenRepository:  repositories.NewAccessTokenRepository(),
		clientRepository:       repositories.NewClientRepository(),
		refreshTokenRepository: repositories.NewRefreshTokenRepository(),
	}
}

func (s *ClientCredentialService) Token(request *requests.TokenRequest, ctx context.Context) (*responses.TokenResponse, error) {
	var clientData models.Client
	err := s.clientRepository.GetByClientId(&clientData, request.ClientId, ctx)

	if err != nil {
		return nil, &schemas.ResponseApiError{
			Status:  schemas.ApiErrorNotFound,
			Message: "Client not found!",
		}
	}

	// Generate access token
	tokenString, err := facades.GenerateToken(clientData.Secret, *clientData.ClientId, 0, 2)

	if err != nil {
		return nil, &schemas.ResponseApiError{
			Status:  schemas.ApiErrorUnprocessAble,
			Message: "Something went wrong when generating access token",
		}
	}

	token, err := facades.ParseToken(tokenString, clientData.Secret)

	if err != nil {
		return nil, &schemas.ResponseApiError{
			Status:  schemas.ApiErrorUnprocessAble,
			Message: err.Error(),
		}
	}

	tokenExpired, _ := token.Claims.GetExpirationTime()

	errSaveToken := s.accessTokenRepository.Create(&models.AccessToken{
		Token:      tokenString,
		UserId:     0,
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

	refreshTokenString, err := facades.GenerateToken(clientData.Secret, *clientData.ClientId, 0, 4)

	if err != nil {
		return nil, &schemas.ResponseApiError{
			Status:  schemas.ApiErrorUnprocessAble,
			Message: "Something went wrong when generating refresh token",
		}
	}

	refreshTokenExpired, _ := token.Claims.GetExpirationTime()

	errSaveRefreshToken := s.refreshTokenRepository.Create(&models.RefreshToken{
		Token:      refreshTokenString,
		UserId:     0,
		ClientId:   clientData.ID,
		ExpiryTime: refreshTokenExpired.Time,
	})

	if errSaveRefreshToken != nil {
		return nil, &schemas.ResponseApiError{
			Status:  schemas.ApiErrorUnprocessAble,
			Message: errSaveRefreshToken.Error(),
		}
	}

	redirectUri := clientData.RedirectUri

	if request.RedirectUri != "" {
		redirectUri = request.RedirectUri
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
		RedirectUri: redirectUri,
	}

	return &res, nil
}

func (s *ClientCredentialService) ValidateToken(request *requests.ValidateTokenRequest) (*responses.ValidateTokenResponse, error) {
	token, err := facades.ParseToken(request.Token, request.Secret)

	if err != nil {
		return nil, err
	}

	tokenExp, _ := token.Claims.GetExpirationTime()
	clientId, _ := facades.GetClientIdFromToken(request.Token, request.Secret)
	userId, _ := facades.GetUserIdFromToken(request.Token, request.Secret)
	res := &responses.ValidateTokenResponse{
		Active:   true,
		Exp:      *tokenExp,
		ClientId: clientId,
		UserId:   userId,
	}
	return res, nil
}
