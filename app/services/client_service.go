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
)

type ClientService struct {
	clientRepository *repositories.ClientRepository
}

func NewClientService() *ClientService {
	return &ClientService{
		clientRepository: repositories.NewClientRepository(),
	}
}

func (s *ClientService) Create(request *requests.ClientRequest, ctx context.Context) error {

	data := models.Client{
		Name:        request.Name,
		Secret:      request.Secret,
		RedirectUri: request.RedirectUri,
	}

	errCreate := s.clientRepository.Create(&data, ctx)

	if errCreate != nil {
		return &schemas.ResponseApiError{
			Status:  schemas.ApiErrorUnprocessAble,
			Message: "Something went wrong while creating client data",
		}
	}

	return nil
}

func (s *ClientService) List(page, limit, sort string, ctx context.Context) (*utils.Pagination, error) {
	var scan []models.Client
	var resp []responses.ClientDetail
	paginateRequest := utils.PaginationRequest{
		Page:  page,
		Limit: limit,
		Sort:  sort,
	}
	pPage, pLimit, pSort := paginateRequest.SetPagination()

	res, err := s.clientRepository.List(&scan, pPage, pLimit, pSort, ctx)

	if err != nil {
		return nil, &schemas.ResponseApiError{
			Status:  schemas.ApiErrorUnprocessAble,
			Message: err.Error(),
		}
	}

	for _, v := range scan {
		resp = append(resp, responses.ClientDetail{
			Id:          v.ID,
			ClientId:    *v.ClientId,
			Secret:      v.Secret,
			RedirectUri: v.RedirectUri,
		})
	}

	res.Rows = &resp

	return res, nil
}

func (s *ClientService) Detail(clientId uint, ctx context.Context) (res *responses.ClientDetail, err error) {
	var scan responses.ClientDetail
	err = s.clientRepository.Detail(&scan, clientId, ctx)

	if err != nil {
		return nil, &schemas.ResponseApiError{
			Status:  schemas.ApiErrorNotFound,
			Message: err.Error(),
		}
	}
	res = &scan
	return
}

func (s *ClientService) Update(clientId uint, request *requests.ClientRequest, ctx context.Context) error {
	// check client
	var clientData models.Client
	err := s.clientRepository.GetById(&clientData, clientId, ctx)
	if err != nil {
		return &schemas.ResponseApiError{
			Status:  schemas.ApiErrorNotFound,
			Message: "Client not found",
		}
	}

	data := &models.Client{
		Name:        request.Name,
		Secret:      request.Secret,
		RedirectUri: request.RedirectUri,
	}

	err = s.clientRepository.Update(data, clientId, ctx)

	if err != nil {
		return &schemas.ResponseApiError{
			Status:  schemas.ApiErrorUnprocessAble,
			Message: "Something went wrong while updating client data",
		}
	}
	return nil
}

func (s *ClientService) GenerateSecret() string {
	return facades.RandomString(32)
}

func (s *ClientService) Delete(clientId uint, ctx context.Context) error {
	var clientData models.Client
	err := s.clientRepository.GetById(&clientData, clientId, ctx)
	if err != nil {
		return &schemas.ResponseApiError{
			Status:  schemas.ApiErrorNotFound,
			Message: "Client not found",
		}
	}

	err = s.clientRepository.Delete(clientId, ctx)

	if err != nil {
		return &schemas.ResponseApiError{
			Status:  schemas.ApiErrorUnprocessAble,
			Message: "Something went wrong while deleting client data",
		}
	}

	return nil
}
