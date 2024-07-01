package services

import (
	"sso-auth/app/repositories"
	"sso-auth/app/responses"
	"sso-auth/app/utils"
)

type HistoryService struct {
	accessTokenRepository  *repositories.AccessTokenRepository
	refreshTokenRepository *repositories.RefreshTokenRepository
	authCodeRepository     *repositories.AuthCodeRepository
}

func NewHistoryService() *HistoryService {
	return &HistoryService{
		accessTokenRepository:  repositories.NewAccessTokenRepository(),
		refreshTokenRepository: repositories.NewRefreshTokenRepository(),
		authCodeRepository:     repositories.NewAuthCodeRepository(),
	}
}

func (s *HistoryService) TokenHistory(page, limit, sort string) (*utils.Pagination, error) {
	var scan []*responses.HistoryResponse

	paginateRequest := utils.PaginationRequest{
		Page:  page,
		Limit: limit,
		Sort:  sort,
	}
	pPage, pLimit, pSort := paginateRequest.SetPagination()
	res, err := s.accessTokenRepository.GetHistory(&scan, pPage, pLimit, pSort)

	if err != nil {
		return nil, err
	}

	res.Rows = scan

	return res, nil
}

func (s *HistoryService) RefreshTokenHistory(page, limit, sort string) (*utils.Pagination, error) {
	var scan []*responses.HistoryResponse

	paginateRequest := utils.PaginationRequest{
		Page:  page,
		Limit: limit,
		Sort:  sort,
	}
	pPage, pLimit, pSort := paginateRequest.SetPagination()
	res, err := s.refreshTokenRepository.GetHistory(&scan, pPage, pLimit, pSort)

	if err != nil {
		return nil, err
	}

	res.Rows = scan

	return res, nil
}

func (s *HistoryService) AuthCodeHistory(page, limit, sort string) (*utils.Pagination, error) {
	var scan []*responses.HistoryResponse

	paginateRequest := utils.PaginationRequest{
		Page:  page,
		Limit: limit,
		Sort:  sort,
	}
	pPage, pLimit, pSort := paginateRequest.SetPagination()
	res, err := s.authCodeRepository.GetHistory(&scan, pPage, pLimit, pSort)

	if err != nil {
		return nil, err
	}

	res.Rows = scan

	return res, nil
}
