package repositories

import (
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

func (r *AdminRepository) Get(data *models.Admin, email string) error {
	res := r.orm.Where("email = ?", email).First(data)

	return res.Error
}
