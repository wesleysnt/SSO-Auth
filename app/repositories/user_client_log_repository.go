package repositories

import (
	"sso-auth/app/facades"
	"sso-auth/app/models"
	"sso-auth/app/responses"
	"sso-auth/app/utils"

	"gorm.io/gorm"
)

type UserClientLogRepository struct {
	orm *gorm.DB
}

func NewUserClientLogRepository() *UserClientLogRepository {
	return &UserClientLogRepository{orm: facades.Orm()}
}

func (r *UserClientLogRepository) Create(data *models.UserClientLog) error {
	return r.orm.Create(data).Error
}

func (r *UserClientLogRepository) Check(userId, clientId uint) bool {
	var count int64
	r.orm.Model(&models.UserClientLog{}).Where("user_id = ? AND client_id = ?", userId, clientId).Count(&count)
	return count > 0
}

func (r *UserClientLogRepository) List(scan *[]responses.UserClientLogResponse, page, limit int, sort string) (*utils.Pagination, error) {
	var pagination utils.Pagination

	var paginate func(methods *gorm.DB) *gorm.DB

	q := r.orm.Model(&models.UserClientLog{})

	paginate = pagination.SetLimit(limit).SetPage(page).SetSort(sort).Pagination(q)

	preQuery := facades.Orm().Scopes(paginate).Model(&models.UserClientLog{}).
		Joins("INNER JOIN users on user_client_logs.user_id = users.id").
		Joins("INNER JOIN clients on user_client_logs.client_id = clients.id").
		Select("user_client_logs.*", "users.name as user", "clients.client_id as client")

	err := preQuery.Scan(&scan)

	return &pagination, err.Error
}
