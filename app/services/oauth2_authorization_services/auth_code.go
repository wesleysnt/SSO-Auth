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
	clientRepository           *repositories.ClientRepository
	authRepository             *repositories.AuthRepository
	authCodeRepository         *repositories.AuthCodeRepository
	accessTokenRepository      *repositories.AccessTokenRepository
	refreshTokenRepository     *repositories.RefreshTokenRepository
	createUserClientLogService *CreateUserClientLogService
}

func NewAuthCodeService() *AuthCodeService {
	return &AuthCodeService{
		clientRepository:           repositories.NewClientRepository(),
		authRepository:             repositories.NewAuthRepository(),
		authCodeRepository:         repositories.NewAuthCodeRepository(),
		accessTokenRepository:      repositories.NewAccessTokenRepository(),
		refreshTokenRepository:     repositories.NewRefreshTokenRepository(),
		createUserClientLogService: NewCreateUserClientLogService(),
	}
}

func (s *AuthCodeService) Login(request *requests.OAuth2LoginRequest) (any, error) {
	var clientData models.Client
	err := s.clientRepository.GetByClientId(&clientData, request.ClientId)

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

	authCode := facades.RandomString(64)

	authCodeData := models.AuthCode{
		Code:       &authCode,
		ExpiryTime: time.Now().Add(time.Minute * 10),
		ClientId:   clientData.ID,
		UserId:     userData.ID,
	}
	err = s.authCodeRepository.Create(&authCodeData)

	res := responses.LoginResponsesAuthCode{
		Id:    authCodeData.UserId,
		Name:  userData.Name,
		Email: userData.Email,
		AuthCode: responses.AuthCode{
			Code:       *authCodeData.Code,
			ExpiryTime: authCodeData.ExpiryTime.Unix(),
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

func (s *AuthCodeService) Token(request *requests.TokenRequest) (*responses.TokenResponse, error) {
	var authCode models.AuthCode

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

	err = s.authCodeRepository.GetCode(request.AuthCode, request.UserId, clientData.ID, &authCode)

	if err != nil {
		return nil, &schemas.ResponseApiError{
			Status:  schemas.ApiErrorUnprocessAble,
			Message: "authorization code is invalid or unvailable",
		}
	}
	if time.Now().After(authCode.ExpiryTime) {
		return nil, &schemas.ResponseApiError{
			Status:  schemas.ApiErrorBadRequest,
			Message: "authorization code is expired",
		}
	}

	// Generate access token
	tokenString, err := facades.GenerateToken(clientData.Secret, userData.ID, clientData.ID, 2)

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

	refreshTokenString, err := facades.GenerateToken(clientData.Secret, userData.ID, clientData.ID, 4)

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

	err = s.createUserClientLogService.Create(userData.ID, clientData.ID)

	if err != nil {
		return nil, &schemas.ResponseApiError{
			Status:  schemas.ApiErrorUnprocessAble,
			Message: err.Error(),
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
		Scope:       request.AuthCode,
		RedirectUri: redirectUri,
	}

	return &res, nil
}
