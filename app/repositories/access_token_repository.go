package repositories

import (
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

func (r *AccessTokenRepository) Create(data *models.AccessToken) error {
	res := r.orm.Create(&data)
	return res.Error
}

func (r *AccessTokenRepository) GetHistory(scan any, page, limit int, sort string) (*utils.Pagination, error) {
	var pagination utils.Pagination

	var paginate func(methods *gorm.DB) *gorm.DB

	q := r.orm.Model(&models.AccessToken{})

	paginate = pagination.SetLimit(limit).SetPage(page).SetSort(sort).Pagination(q)

	res := facades.Orm().Scopes(paginate).Model(&models.AccessToken{}).Joins("left join users on users.id = access_tokens.user_id").Joins("left join clients on clients.id = access_tokens.client_id").Select("access_tokens.id as id", "access_tokens.created_at as date", "clients.client_id as client", "users.username as user").Scan(scan)

	return &pagination, res.Error
}
