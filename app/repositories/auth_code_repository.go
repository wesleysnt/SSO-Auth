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

func (r *AuthCodeRepository) GetCode(code string, userId, clientId uint, authCode *models.AuthCode) (err error) {
	res := r.orm.Where("code", code).Where("client_id", clientId).Where("user_id", userId).First(&authCode)
	return res.Error
}
