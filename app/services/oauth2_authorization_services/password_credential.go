package oauth2authorizationservices

import (
	"context"
	"sso-auth/app/facades"
	"sso-auth/app/http/requests"
	"sso-auth/app/models"
	"sso-auth/app/repositories"
	"sso-auth/app/responses"
	"sso-auth/app/schemas"
	"sso-auth/app/utils"
)

type PasswordCredetialService struct {
	clientRepository           *repositories.ClientRepository
	authRepository             *repositories.AuthRepository
	accessTokenRepository      *repositories.AccessTokenRepository
	refreshTokenRepository     *repositories.RefreshTokenRepository
	createUserClientLogService *CreateUserClientLogService
}

func NewPasswordCredentialService() *PasswordCredetialService {
	return &PasswordCredetialService{
		clientRepository:           repositories.NewClientRepository(),
		authRepository:             repositories.NewAuthRepository(),
		accessTokenRepository:      repositories.NewAccessTokenRepository(),
		refreshTokenRepository:     repositories.NewRefreshTokenRepository(),
		createUserClientLogService: NewCreateUserClientLogService(),
	}
}

func (s *PasswordCredetialService) Login(request *requests.OAuth2LoginRequest, ctx context.Context) (any, error) {
	var clientData models.Client
	err := s.clientRepository.GetByClientId(&clientData, request.ClientId, ctx)

	if err != nil {
		return nil, &schemas.ResponseApiError{
			Status:  schemas.ApiErrorNotFound,
			Message: "Client not found",
		}
	}

	var userData models.User

	err = s.authRepository.GetUser(&userData, request.Email)
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

	// Generate access token
	tokenString, err := facades.GenerateToken(clientData.Secret, *clientData.ClientId, userData.ID, 2)

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
		UserId:     userData.ID,
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

	refreshTokenString, err := facades.GenerateToken(clientData.Secret, *clientData.ClientId, userData.ID, 4)

	if err != nil {
		return nil, &schemas.ResponseApiError{
			Status:  schemas.ApiErrorUnprocessAble,
			Message: "Something went wrong when generating refresh token",
		}
	}

	refreshTokenExpired, _ := token.Claims.GetExpirationTime()

	errSaveRefreshToken := s.refreshTokenRepository.Create(&models.RefreshToken{
		Token:      refreshTokenString,
		UserId:     userData.ID,
		ClientId:   clientData.ID,
		ExpiryTime: refreshTokenExpired.Time,
	})

	if errSaveRefreshToken != nil {
		return nil, &schemas.ResponseApiError{
			Status:  schemas.ApiErrorUnprocessAble,
			Message: errSaveRefreshToken.Error(),
		}
	}

	res := responses.LoginResponses{
		Id:    userData.ID,
		Name:  userData.Name,
		Email: userData.Email,
		AccessToken: responses.AccessToken{
			Token:      tokenString,
			ExpiryTime: tokenExpired,
		},
		RefreshToken: responses.RefreshToken{
			Token:      refreshTokenString,
			ExpiryTime: refreshTokenExpired,
		},
	}

	err = s.createUserClientLogService.Create(userData.ID, clientData.ID)

	if err != nil {
		return nil, &schemas.ResponseApiError{
			Status:  schemas.ApiErrorUnprocessAble,
			Message: err.Error(),
		}
	}

	return &res, nil
}

func (s *PasswordCredetialService) ValidateToken(request *requests.ValidateTokenRequest) (*responses.ValidateTokenResponse, error) {
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
