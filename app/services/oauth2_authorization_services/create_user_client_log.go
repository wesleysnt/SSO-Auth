package oauth2authorizationservices

import (
	"sso-auth/app/models"
	"sso-auth/app/repositories"
)

type CreateUserClientLogService struct {
	createUserClientLogRepository *repositories.UserClientLogRepository
}

func NewCreateUserClientLogService() *CreateUserClientLogService {
	return &CreateUserClientLogService{
		createUserClientLogRepository: repositories.NewUserClientLogRepository(),
	}
}

func (s *CreateUserClientLogService) Create(userId, clientId uint) error {
	err := s.createUserClientLogRepository.Check(userId, clientId)
	if err {
		return nil
	}
	return s.createUserClientLogRepository.Create(&models.UserClientLog{
		UserId:   userId,
		ClientId: clientId,
	})
}
