package repositories

import (
	"sso-auth/app/facades"
	"sso-auth/app/models"
	"sso-auth/app/responses"
	"time"

	"gorm.io/gorm"
)

type DashboardRepository struct {
	orm *gorm.DB
}

func NewDashboardRepository() *DashboardRepository {
	return &DashboardRepository{
		orm: facades.Orm(),
	}
}

func (r *DashboardRepository) GetTotalActiveToken() int64 {
	var total int64
	r.orm.Model(&models.AccessToken{}).Where("expiry_time >= ?", time.Now()).Count(&total)

	return total
}

func (r *DashboardRepository) GetTotalUser() int64 {
	var total int64
	r.orm.Model(&models.User{}).Count(&total)

	return total
}

func (r *DashboardRepository) GetLatestLog() []*responses.LatestLog {
	var data []*responses.LatestLog
	r.orm.Model(&models.UserClientLog{}).Joins("INNER JOIN users ON users.id = user_client_logs.user_id").Joins("INNER JOIN clients ON clients.id = user_client_logs.client_id").Order("user_client_logs.created_at desc").Select("user_client_logs.user_id", "user_client_logs.client_id", "clients.name as client_name", "users.name as user_name", "user_client_logs.created_at as log_date").Limit(5).Scan(&data)

	return data
}
