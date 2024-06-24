package repositories

import (
	"sso-auth/app/facades"
	"sso-auth/app/models"

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
