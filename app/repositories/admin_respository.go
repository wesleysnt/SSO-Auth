package repositories

import (
	"context"
	"sso-auth/app/facades"
	"sso-auth/app/models"

	"gorm.io/gorm"
)

type AdminRepository struct {
	orm *gorm.DB
}

func NewAdminRepository() *AdminRepository {
	return &AdminRepository{
		orm: facades.Orm(),
	}
}

func (r *AdminRepository) Get(data *models.Admin, email string, ctx context.Context) error {
	res := r.orm.WithContext(ctx).Where("email = ?", email).First(data)

	return res.Error
}
