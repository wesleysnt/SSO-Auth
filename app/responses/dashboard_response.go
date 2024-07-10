package responses

import (
	"github.com/goravel/framework/support/carbon"
)

type DashboardResponse struct {
	TotalActiveToken int64        `json:"total_active_token"`
	TotalUser        int64        `json:"total_user"`
	LatestLog        []*LatestLog `json:"latest_log"`
}

type LatestLog struct {
	UserId     uint            `json:"user_id"`
	UserName   string          `json:"user_name"`
	ClientId   uint            `json:"client_id"`
	ClientName string          `json:"client_name"`
	LogDate    carbon.DateTime `json:"log_date"`
}
