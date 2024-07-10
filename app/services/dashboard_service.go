package services

import (
	"sso-auth/app/repositories"
	"sso-auth/app/responses"
)

type DashboardService struct {
	dashboardRepositry *repositories.DashboardRepository
}

func NewDashboardService() *DashboardService {
	return &DashboardService{
		dashboardRepositry: repositories.NewDashboardRepository(),
	}
}

func (s *DashboardService) Get() (*responses.DashboardResponse, error) {
	activeToken := s.dashboardRepositry.GetTotalActiveToken()
	totalUser := s.dashboardRepositry.GetTotalUser()
	latestLog := s.dashboardRepositry.GetLatestLog()

	return &responses.DashboardResponse{
		TotalActiveToken: activeToken,
		TotalUser:        totalUser,
		LatestLog:        latestLog,
	}, nil
}
