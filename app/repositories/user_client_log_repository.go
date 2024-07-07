package repositories

import (
	"sso-auth/app/facades"
	"sso-auth/app/models"

	"gorm.io/gorm"
)

type UserClientLogRepository struct {
	orm *gorm.DB
}

func NewUserClientLogRepository() *UserClientLogRepository {
	return &UserClientLogRepository{orm: facades.Orm()}
}

func (r *UserClientLogRepository) Create(data *models.UserClientLog) error {
	return r.orm.Create(data).Error
}

func (r *UserClientLogRepository) Check(userId, clientId uint) bool {
	var count int64
	r.orm.Model(&models.UserClientLog{}).Where("user_id = ? AND client_id = ?", userId, clientId).Count(&count)
	return count > 0
}
