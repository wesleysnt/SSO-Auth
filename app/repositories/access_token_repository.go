package repositories

import (
	"sso-auth/app/facades"
	"sso-auth/app/models"

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
