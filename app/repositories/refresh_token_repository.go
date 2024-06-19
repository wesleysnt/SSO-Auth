package repositories

import (
	"sso-auth/app/facades"
	"sso-auth/app/models"

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

func (r *RefreshTokenRepository) Create(data *models.RefreshToken) error {
	res := r.orm.Create(&data)
	return res.Error
}
