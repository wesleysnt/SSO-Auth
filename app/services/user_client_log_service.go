package services

import (
	"context"
	"sso-auth/app/repositories"
	"sso-auth/app/responses"
	"sso-auth/app/schemas"
	"sso-auth/app/utils"
)

type UserClientLogService struct {
	userClientLogRepository *repositories.UserClientLogRepository
}

func NewUserClientLogService() *UserClientLogService {
	return &UserClientLogService{
		userClientLogRepository: repositories.NewUserClientLogRepository(),
	}
}

func (s *UserClientLogService) List(page, limit, sort string, ctx context.Context) (*utils.Pagination, error) {
	var scan []responses.UserClientLogResponse
	paginateRequest := utils.PaginationRequest{
		Page:  page,
		Limit: limit,
		Sort:  sort,
	}
	pPage, pLimit, pSort := paginateRequest.SetPagination()

	res, err := s.userClientLogRepository.List(&scan, pPage, pLimit, pSort, ctx)

	if err != nil {
		return nil, &schemas.ResponseApiError{
			Status:  schemas.ApiErrorUnprocessAble,
			Message: err.Error(),
		}
	}

	res.Rows = &scan

	return res, nil
}
