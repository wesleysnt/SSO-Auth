package services

import (
	"sso-auth/app/facades"
	"sso-auth/app/http/requests"
	"sso-auth/app/models"
	"sso-auth/app/repositories"
	"sso-auth/app/responses"
	"sso-auth/app/schemas"
	"sso-auth/app/utils"
)

type AuthService struct {
	adminRepository *repositories.AdminRepository
}

func NewAuthService() *AuthService {
	return &AuthService{
		adminRepository: repositories.NewAdminRepository(),
	}
}

func (s *AuthService) Login(request *requests.LoginRequest) (*responses.AdminLoginResponses, error) {
	// check admin email
	var adminData models.Admin
	err := s.adminRepository.Get(&adminData, request.Email)
	if err != nil {
		return nil, &schemas.ResponseApiError{
			Status:  schemas.ApiErrorForbidden,
			Message: "Email or Password incorrect",
		}
	}

	auth := utils.CheckPasswordHash(request.Password, adminData.Password)

	if !auth {
		return nil, &schemas.ResponseApiError{
			Status:  schemas.ApiErrorForbidden,
			Message: "Email or Password incorrect",
		}
	}

	// Generate access token
	tokenString, err := facades.GenerateToken("", adminData.ID, 0, 2)

	if err != nil {
		return nil, &schemas.ResponseApiError{
			Status:  schemas.ApiErrorUnprocessAble,
			Message: "Something went wrong when generating access token",
		}
	}

	token, err := facades.ParseToken(tokenString, "")

	if err != nil {
		return nil, &schemas.ResponseApiError{
			Status:  schemas.ApiErrorUnprocessAble,
			Message: err.Error(),
		}
	}

	tokenExpired, _ := token.Claims.GetExpirationTime()

	res := responses.AdminLoginResponses{
		Id:    adminData.ID,
		Email: adminData.Email,
		AccessToken: responses.AccessToken{
			Token:      tokenString,
			ExpiryTime: tokenExpired,
		},
	}

	return &res, nil

}
