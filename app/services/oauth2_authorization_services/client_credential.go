package oauth2authorizationservices

import (
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

func (s *ClientCredentialService) Token(request *requests.TokenRequest) (*responses.TokenResponse, error) {
	var clientData models.Client
	err := s.clientRepository.GetByClientId(&clientData, request.ClientId)

	if err != nil {
		return nil, &schemas.ResponseApiError{
			Status:  schemas.ApiErrorNotFound,
			Message: "Client not found!",
		}
	}

	// Generate access token
	tokenString, err := facades.GenerateToken(clientData.Secret, 0, clientData.ID, 2)

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

	refreshTokenString, err := facades.GenerateToken(clientData.Secret, 0, clientData.ID, 4)

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
