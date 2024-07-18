package repositories

import (
	"context"
	"sso-auth/app/facades"
	"sso-auth/app/models"
	"sso-auth/app/utils"

	"gorm.io/gorm"
)

type AccessTokenRepository struct {
	orm *gorm.DB
}

func NewAccessTokenRepository() *AccessTokenRepository {
	return &AccessTokenRepository{
		orm: facades.Orm(),
	}
}

func (r *AccessTokenRepository) Create(data *models.AccessToken, ctx context.Context) error {
	res := r.orm.WithContext(ctx).Create(&data)
	return res.Error
}

func (r *AccessTokenRepository) GetByToken(token string, ctx context.Context) (*models.AccessToken, error) {
	data := models.AccessToken{}
	res := r.orm.WithContext(ctx).Where("token = ?", token).First(&data)
	return &data, res.Error
}

func (r *AccessTokenRepository) GetByTokenAndClient(token string, clientId uint, ctx context.Context) (*models.AccessToken, error) {
	data := models.AccessToken{}
	res := r.orm.WithContext(ctx).Where("token = ?", token).Where("client_id = ?", clientId).First(&data)
	return &data, res.Error
}

func (r *AccessTokenRepository) GetHistory(scan any, page, limit int, sort string, ctx context.Context) (*utils.Pagination, error) {
	var pagination utils.Pagination

	var paginate func(methods *gorm.DB) *gorm.DB

	q := r.orm.WithContext(ctx).Model(&models.AccessToken{})

	paginate = pagination.SetLimit(limit).SetPage(page).SetSort(sort).Pagination(q)

	res := facades.Orm().WithContext(ctx).Scopes(paginate).Model(&models.AccessToken{}).Joins("left join users on users.id = access_tokens.user_id").Joins("left join clients on clients.id = access_tokens.client_id").Select("access_tokens.id as id", "access_tokens.created_at as date", "clients.name as client", "users.name as user").Scan(scan)

	return &pagination, res.Error
}
