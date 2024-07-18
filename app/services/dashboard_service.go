package services

import (
	"context"
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

func (s *DashboardService) Get(ctx context.Context) (*responses.DashboardResponse, error) {
	activeToken := s.dashboardRepositry.GetTotalActiveToken(ctx)
	totalUser := s.dashboardRepositry.GetTotalUser(ctx)
	latestLog := s.dashboardRepositry.GetLatestLog(ctx)

	return &responses.DashboardResponse{
		TotalActiveToken: activeToken,
		TotalUser:        totalUser,
		LatestLog:        latestLog,
	}, nil
}
