package repositories

import (
	"sso-auth/app/facades"
	"sso-auth/app/models"

	"gorm.io/gorm"
)

type ForgotPasswordTokenRepository struct {
	orm *gorm.DB
}

func NewForgotPasswordTokenRepository() *ForgotPasswordTokenRepository {
	return &ForgotPasswordTokenRepository{orm: facades.Orm()}
}

func (r *ForgotPasswordTokenRepository) Create(token *models.ForgotPasswordToken) error {
	return r.orm.Create(&token).Error
}

func (r *ForgotPasswordTokenRepository) FindByToken(token string, forgotPasswordToken *models.ForgotPasswordToken) error {
	err := r.orm.Where("token = ?", token).First(&forgotPasswordToken).Error
	return err
}

func (r *ForgotPasswordTokenRepository) UpdateIsUsed(token string) error {
	err := r.orm.Model(&models.ForgotPasswordToken{}).Where("token", token).UpdateColumn("is_used", true).Error
	return err
}
