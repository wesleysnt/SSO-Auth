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
	codeChallengeRepository    *repositories.CodeChallengeRepository
}

func NewAuthCodeService() *AuthCodeService {
	return &AuthCodeService{
		clientRepository:           repositories.NewClientRepository(),
		authRepository:             repositories.NewAuthRepository(),
		authCodeRepository:         repositories.NewAuthCodeRepository(),
		accessTokenRepository:      repositories.NewAccessTokenRepository(),
		refreshTokenRepository:     repositories.NewRefreshTokenRepository(),
		createUserClientLogService: NewCreateUserClientLogService(),
		codeChallengeRepository:    repositories.NewCodeChallengeRepository(),
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

	if err != nil {
		return nil, &schemas.ResponseApiError{
			Status:  schemas.ApiErrorUnprocessAble,
			Message: err.Error(),
		}
	}

	codeChallengeData := models.CodeChallenge{
		Code:     request.CodeChallenge,
		ClientId: clientData.ID,
		Method:   "s256",
	}

	err = s.codeChallengeRepository.Create(&codeChallengeData)

	if err != nil {
		return nil, &schemas.ResponseApiError{
			Status:  schemas.ApiErrorUnprocessAble,
			Message: err.Error(),
		}
	}

	res := responses.LoginResponsesAuthCode{
		Id:    authCodeData.UserId,
		Name:  userData.Name,
		Email: userData.Email,
		AuthCode: responses.AuthCode{
			Code:       *authCodeData.Code,
			ExpiryTime: authCodeData.ExpiryTime.Unix(),
		},
		CodeChallengeUniqueCode: *codeChallengeData.UniqueCode,
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

	var codeChallengeData models.CodeChallenge
	err = s.codeChallengeRepository.GetChallenge(request.CodeChallengeUniqueCode, &codeChallengeData)

	if err != nil || !utils.VerifyCode(request.CodeVerifier, codeChallengeData.Code) || codeChallengeData.ClientId != clientData.ID {
		return nil, &schemas.ResponseApiError{
			Status:  schemas.ApiErrorBadRequest,
			Message: "Invalid code verifier",
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
		Scope:       request.Scope,
		RedirectUri: redirectUri,
	}

	return &res, nil
}

func (s *AuthCodeService) ValidateToken(request *requests.ValidateTokenRequest) (*responses.ValidateTokenResponse, error) {
	var codeChallengeData models.CodeChallenge
	s.codeChallengeRepository.GetChallenge(request.CodeChallengeUniqueCode, &codeChallengeData)
	if !utils.VerifyCode(request.CodeVerifier, codeChallengeData.Code) {
		return nil, &schemas.ResponseApiError{
			Status:  schemas.ApiErrorBadRequest,
			Message: "Invalid code verifier",
		}
	}

	tokenData, err := s.accessTokenRepository.GetByToken(request.Token)
	if err != nil {
		return nil, &schemas.ResponseApiError{
			Status:  schemas.ApiErrorUnauthorized,
			Message: "Invalid token",
		}
	}
	var clientData models.Client
	s.clientRepository.GetById(&clientData, tokenData.ClientId)
	// Verify secret
	token, err := facades.ParseToken(request.Token, clientData.Secret)

	if err != nil {
		return nil, err
	}

	tokenExp, _ := token.Claims.GetExpirationTime()
	clientId, _ := facades.GetClientIdFromToken(request.Token, clientData.Secret)
	userId, _ := facades.GetUserIdFromToken(request.Token, clientData.Secret)
	res := &responses.ValidateTokenResponse{
		Active:   true,
		Exp:      *tokenExp,
		ClientId: clientId,
		UserId:   userId,
	}
	return res, nil
}
