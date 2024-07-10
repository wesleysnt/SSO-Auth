package services

import (
	"sso-auth/app/models"
	"sso-auth/app/repositories"
	"sso-auth/app/responses"
	"sso-auth/app/schemas"
	"sso-auth/app/utils"
)

type UserService struct {
	userRepository *repositories.AuthRepository
}

func NewUserService() *UserService {
	return &UserService{
		userRepository: repositories.NewAuthRepository(),
	}
}

func (s *UserService) List(page, limit, sort string) (*utils.Pagination, error) {
	var scan []models.User
	var resp []responses.UserResponses
	paginateRequest := utils.PaginationRequest{
		Page:  page,
		Limit: limit,
		Sort:  sort,
	}
	pPage, pLimit, pSort := paginateRequest.SetPagination()

	res, err := s.userRepository.List(&scan, pPage, pLimit, pSort)

	if err != nil {
		return nil, err
	}

	for _, v := range scan {
		resp = append(resp, responses.UserResponses{
			Id:       v.ID,
			Name:     v.Name,
			Email:    v.Email,
			Phone:    v.Phone,
			IsActive: v.IsActive,
		})
	}

	res.Rows = &resp

	return res, nil
}

func (s *UserService) GetById(id uint) (*responses.UserResponses, error) {
	var scan models.User
	var resp responses.UserResponses

	err := s.userRepository.GetById(&scan, id)

	if err != nil {
		return nil, &schemas.ResponseApiError{
			Status:  schemas.ApiErrorNotFound,
			Message: "User not found",
		}
	}

	resp = responses.UserResponses{
		Id:       scan.ID,
		Name:     scan.Name,
		Email:    scan.Email,
		Phone:    scan.Phone,
		IsActive: scan.IsActive,
	}

	return &resp, nil
}
