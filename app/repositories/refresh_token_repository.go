package repositories

import (
	"context"
	"sso-auth/app/facades"
	"sso-auth/app/models"
	"sso-auth/app/utils"

	"gorm.io/gorm"
)

type RefreshTokenRepository struct {
	orm *gorm.DB
}

func NewRefreshTokenRepository() *RefreshTokenRepository {
	return &RefreshTokenRepository{
		orm: facades.Orm(),
	}
}

func (r *RefreshTokenRepository) Create(data *models.RefreshToken, ctx context.Context) error {
	res := r.orm.WithContext(ctx).Create(&data)
	return res.Error
}

func (r *RefreshTokenRepository) GetHistory(scan any, page, limit int, sort string, ctx context.Context) (*utils.Pagination, error) {
	var pagination utils.Pagination

	var paginate func(methods *gorm.DB) *gorm.DB

	q := r.orm.WithContext(ctx).Model(&models.RefreshToken{})

	paginate = pagination.SetLimit(limit).SetPage(page).SetSort(sort).Pagination(q)

	res := facades.Orm().Scopes(paginate).Model(&models.RefreshToken{}).Joins("left join users on users.id = refresh_tokens.user_id").Joins("left join clients on clients.id = refresh_tokens.client_id").Select("refresh_tokens.id as id", "refresh_tokens.created_at as date", "clients.name as client", "users.name as user").Scan(scan)

	return &pagination, res.Error
}

func (r *RefreshTokenRepository) Check(token string, clientId, userId uint, ctx context.Context) error {
	var refreshToken models.RefreshToken
	res := r.orm.WithContext(ctx).Where("token = ? and client_id = ? and user_id = ?", token, clientId, userId).First(&refreshToken)
	if res.Error != nil {
		return res.Error
	}
	return nil
}
