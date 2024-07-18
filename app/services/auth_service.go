package services

import (
	"context"
	"sso-auth/app/facades"
	"sso-auth/app/http/requests"
	"sso-auth/app/models"
	"sso-auth/app/repositories"
	"sso-auth/app/responses"
	"sso-auth/app/schemas"
	"sso-auth/app/utils"
	"sso-auth/app/utils/otel"

	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"
)

type AuthService struct {
	adminRepository *repositories.AdminRepository
}

func NewAuthService() *AuthService {
	return &AuthService{
		adminRepository: repositories.NewAdminRepository(),
	}
}

func (s *AuthService) Login(request *requests.LoginRequest, ctx context.Context) (*responses.AdminLoginResponses, error) {

	ctxLogin, span := otel.StartNewTrace(ctx, "Start Login")
	defer otel.EndSpan(span)

	// check admin email
	otel.AddEvent(span, "Checking admin email", trace.WithAttributes(attribute.String("email", request.Email)))
	var adminData models.Admin
	err := s.adminRepository.Get(&adminData, request.Email, ctxLogin)
	if err != nil {
		return nil, &schemas.ResponseApiError{
			Status:  schemas.ApiErrorForbidden,
			Message: "Email or Password incorrect",
		}
	}

	otel.AddEvent(span, "check password")
	auth := utils.CheckPasswordHash(request.Password, adminData.Password, ctx)

	if !auth {
		return nil, &schemas.ResponseApiError{
			Status:  schemas.ApiErrorForbidden,
			Message: "Email or Password incorrect",
		}
	}

	// Generate access token
	otel.AddEvent(span, "Generating access token")
	tokenString, err := facades.GenerateToken("", "", adminData.ID, 2, ctx)

	if err != nil {
		return nil, &schemas.ResponseApiError{
			Status:  schemas.ApiErrorUnprocessAble,
			Message: "Something went wrong when generating access token",
		}
	}

	token, err := facades.ParseToken(tokenString, "", ctx)

	if err != nil {
		return nil, &schemas.ResponseApiError{
			Status:  schemas.ApiErrorUnprocessAble,
			Message: err.Error(),
		}
	}

	tokenExpired, _ := token.Claims.GetExpirationTime()

	otel.AddEvent(span, "Generating response")
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
