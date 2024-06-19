package repositories

import (
	"sso-auth/app/facades"
	"sso-auth/app/models"

	"gorm.io/gorm"
)

type AuthRepository struct {
	orm *gorm.DB
}

func NewAuthRepository() *AuthRepository {
	return &AuthRepository{
		orm: facades.Orm(),
	}
}

func (r *AuthRepository) CreateUser(data *models.User) error {
	res := r.orm.Create(&data)
	return res.Error
}

func (r *AuthRepository) GetUser(data *models.User, username string) error {
	res := r.orm.Where("username = ?", username).First(&data)
	return res.Error
}
