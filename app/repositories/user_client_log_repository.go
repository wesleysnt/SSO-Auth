package repositories

import (
	"context"
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

func (r *UserClientLogRepository) Create(data *models.UserClientLog, ctx context.Context) error {
	return r.orm.WithContext(ctx).Create(data).Error
}

func (r *UserClientLogRepository) Check(userId, clientId uint, ctx context.Context) bool {
	var count int64
	r.orm.WithContext(ctx).Model(&models.UserClientLog{}).Where("user_id = ? AND client_id = ?", userId, clientId).Count(&count)
	return count > 0
}

func (r *UserClientLogRepository) List(scan *[]responses.UserClientLogResponse, page, limit int, sort string, ctx context.Context) (*utils.Pagination, error) {
	var pagination utils.Pagination

	var paginate func(methods *gorm.DB) *gorm.DB

	q := r.orm.WithContext(ctx).Model(&models.UserClientLog{})

	paginate = pagination.SetLimit(limit).SetPage(page).SetSort(sort).Pagination(q)

	preQuery := facades.Orm().Scopes(paginate).Model(&models.UserClientLog{}).
		Joins("INNER JOIN users on user_client_logs.user_id = users.id").
		Joins("INNER JOIN clients on user_client_logs.client_id = clients.id").
		Select("user_client_logs.*", "users.name as user", "clients.name as client").Order("user_client_logs.created_at desc")

	err := preQuery.Scan(&scan)

	return &pagination, err.Error
}
