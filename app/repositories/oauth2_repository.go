package repositories

import (
	"context"
	"sso-auth/app/facades"
	"sso-auth/app/models"
	"sso-auth/app/utils"

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

func (r *AuthRepository) CreateUser(data *models.User, ctx context.Context) error {
	res := r.orm.WithContext(ctx).Create(&data)
	return res.Error
}

func (r *AuthRepository) CheckEmailExists(email string, data *models.User, ctx context.Context) error {
	res := r.orm.WithContext(ctx).Where("email = ?", email).First(&data)
	return res.Error
}

func (r *AuthRepository) GetUser(data *models.User, username string, ctx context.Context) error {
	res := r.orm.WithContext(ctx).Where("email = ?", username).First(&data)
	return res.Error
}

func (r *AuthRepository) GetById(data *models.User, id uint, ctx context.Context) error {
	res := r.orm.WithContext(ctx).Where("id = ?", id).First(&data)
	return res.Error
}

func (r *AuthRepository) List(scan *[]models.User, page, limit int, sort string, ctx context.Context) (*utils.Pagination, error) {
	var pagination utils.Pagination

	var paginate func(methods *gorm.DB) *gorm.DB

	q := r.orm.WithContext(ctx).Model(&models.User{})

	paginate = pagination.SetLimit(limit).SetPage(page).SetSort(sort).Pagination(q)

	preQuery := facades.Orm().WithContext(ctx).Scopes(paginate).Model(&models.User{})

	err := preQuery.Scan(&scan)

	return &pagination, err.Error
}

func (r *AuthRepository) UpdateWithTx(tx *gorm.DB, data *models.User, id uint, ctx context.Context) error {
	res := tx.WithContext(ctx).Model(&models.User{}).Where("id = ?", id).Updates(data)
	return res.Error
}

func (r *AuthRepository) UpdatePassword(userId uint, password string, ctx context.Context) error {
	res := r.orm.WithContext(ctx).Where("id = ?", userId).Updates(&models.User{Password: password})
	return res.Error
}
