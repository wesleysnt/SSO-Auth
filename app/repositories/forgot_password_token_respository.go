package repositories

import (
	"context"
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

func (r *ForgotPasswordTokenRepository) Create(token *models.ForgotPasswordToken, ctx context.Context) error {
	return r.orm.WithContext(ctx).Create(&token).Error
}

func (r *ForgotPasswordTokenRepository) FindByToken(token string, forgotPasswordToken *models.ForgotPasswordToken, ctx context.Context) error {
	err := r.orm.WithContext(ctx).Where("token = ?", token).First(&forgotPasswordToken).Error
	return err
}

func (r *ForgotPasswordTokenRepository) UpdateIsUsed(token string, ctx context.Context) error {
	err := r.orm.WithContext(ctx).Model(&models.ForgotPasswordToken{}).Where("token", token).UpdateColumn("is_used", true).Error
	return err
}
