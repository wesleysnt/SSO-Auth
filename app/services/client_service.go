package services

import (
	"sso-auth/app/facades"
	"sso-auth/app/http/requests"
	"sso-auth/app/models"
	"sso-auth/app/repositories"
	"sso-auth/app/schemas"
)

type ClientService struct {
	clientRepository *repositories.ClientRepository
}

func NewClientService() *ClientService {
	return &ClientService{
		clientRepository: repositories.NewClientRepository(),
	}
}

func (s *ClientService) Create(request *requests.ClientRequest) error {
	err := s.clientRepository.CheckClientId(request.ClientId)

	if err == nil {
		return &schemas.ResponseApiError{
			Status:  schemas.ApiErrorNotFound,
			Message: "This client id is already taken",
		}
	}

	secret := facades.RandomString(32)

	data := models.Client{
		ClientId:    request.ClientId,
		Secret:      secret,
		RedirectUri: request.RedirectUri,
	}

	errCreate := s.clientRepository.Create(&data)

	if errCreate != nil {
		return &schemas.ResponseApiError{
			Status:  schemas.ApiErrorUnprocessAble,
			Message: "Something went wrong while creating client data",
		}
	}

	return nil

}
