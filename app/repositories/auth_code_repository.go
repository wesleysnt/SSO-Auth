package repositories

import (
	"sso-auth/app/facades"
	"sso-auth/app/models"
	"sso-auth/app/utils"

	"gorm.io/gorm"
)

type AuthCodeRepository struct {
	orm *gorm.DB
}

func NewAuthCodeRepository() *AuthCodeRepository {
	return &AuthCodeRepository{
		orm: facades.Orm(),
	}
}

func (r *AuthCodeRepository) Create(data *models.AuthCode) error {
	res := r.orm.Create(&data)
	return res.Error
}

func (r *AuthCodeRepository) GetCode(code string, userId, clientId uint, authCode *models.AuthCode) (err error) {
	res := r.orm.Where("code", code).Where("client_id", clientId).Where("user_id", userId).First(&authCode)
	return res.Error
}

func (r *AuthCodeRepository) GetHistory(scan any, page, limit int, sort string) (*utils.Pagination, error) {
	var pagination utils.Pagination

	var paginate func(methods *gorm.DB) *gorm.DB

	q := r.orm.Model(&models.AuthCode{})

	paginate = pagination.SetLimit(limit).SetPage(page).SetSort(sort).Pagination(q)

	res := facades.Orm().Scopes(paginate).Model(&models.AuthCode{}).Joins("left join users on users.id = auth_codes.user_id").Joins("left join clients on clients.id = auth_codes.client_id").Select("auth_codes.id as id", "auth_codes.created_at as date", "clients.client_id as client", "users.username as user").Scan(scan)

	return &pagination, res.Error
}
