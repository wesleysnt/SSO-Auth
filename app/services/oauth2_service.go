package services

import (
	"errors"
	"sso-auth/app/facades"
	"sso-auth/app/http/requests"
	"sso-auth/app/models"
	"sso-auth/app/repositories"
	"sso-auth/app/responses"
	"sso-auth/app/schemas"
	"sso-auth/app/utils"
)

type AuthService struct {
	authRepository        *repositories.AuthRepository
	clientRepository      *repositories.ClientRepository
	accessTokenRepository *repositories.AccessTokenRepository
}

func NewAuthService() *AuthService {
	return &AuthService{
		authRepository:        repositories.NewAuthRepository(),
		clientRepository:      repositories.NewClientRepository(),
		accessTokenRepository: repositories.NewAccessTokenRepository(),
	}
}

func (s *AuthService) Login(grantType, redirectUri string, request *requests.LoginRequest) (*responses.LoginResponses, error) {

	if grantType == string(requests.GrantTypePasswordCredential) {
		var clientData models.Client
		err := s.clientRepository.GetById(&clientData, request.ClientId)

		if err != nil {
			return nil, &schemas.ResponseApiError{
				Status:  schemas.ApiErrorNotFound,
				Message: "Client not found",
			}
		}

		var userData models.User

		err = s.authRepository.GetUser(&userData, request.Username)
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

		tokenString, err := facades.GenerateToken(clientData.Secret, userData.ID)

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
				Message: err.Error(),
			}
		}

		res := responses.LoginResponses{
			Id:       userData.ID,
			Username: userData.Username,
			Email:    userData.Email,
			AccessToken: responses.AccessToken{
				Token:      tokenString,
				ExpiryTime: tokenExpired,
			},
		}

		return &res, nil

	}

	return nil, errors.New("aslkjdaosjd")
}

func (s *AuthService) Register(data *requests.AuthRequest) (err error) {

	user := models.User{
		Username: data.Username,
		Password: data.Password,
		Email:    data.Email,
	}

	err = s.authRepository.CreateUser(&user)

	return
}
