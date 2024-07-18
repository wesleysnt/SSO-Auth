package repositories

import (
	"context"
	"sso-auth/app/facades"
	"sso-auth/app/models"
	"sso-auth/app/responses"
	"sso-auth/app/utils"

	"gorm.io/gorm"
)

type ClientRepository struct {
	orm *gorm.DB
}

func NewClientRepository() *ClientRepository {
	return &ClientRepository{
		orm: facades.Orm(),
	}
}

func (r *ClientRepository) Create(data *models.Client, ctx context.Context) error {
	res := r.orm.Create(&data)
	return res.Error
}
func (r *ClientRepository) GetById(data *models.Client, id uint, ctx context.Context) error {
	res := r.orm.First(&data, id)
	return res.Error
}

func (r *ClientRepository) GetByClientId(data *models.Client, clientId string, ctx context.Context) error {
	res := r.orm.Where("client_id = ?", clientId).First(&data)
	return res.Error
}

func (r *ClientRepository) CheckClientId(clientId string, ctx context.Context) error {
	var data *models.Client
	res := r.orm.Where("client_id = ?", clientId).First(&data)
	return res.Error
}

func (r *ClientRepository) List(scan *[]models.Client, page, limit int, sort string, ctx context.Context) (*utils.Pagination, error) {
	var pagination utils.Pagination

	var paginate func(methods *gorm.DB) *gorm.DB

	q := r.orm.WithContext(ctx).Model(&models.Client{})

	paginate = pagination.SetLimit(limit).SetPage(page).SetSort(sort).Pagination(q)

	preQuery := facades.Orm().WithContext(ctx).Scopes(paginate).Model(&models.Client{})

	err := preQuery.Scan(&scan)

	return &pagination, err.Error
}

func (r *ClientRepository) Detail(scan *responses.ClientDetail, clientId uint, ctx context.Context) error {
	db := r.orm.First(&models.Client{}, clientId).Scan(scan)
	return db.Error
}

func (r *ClientRepository) Update(data *models.Client, clientId uint, ctx context.Context) error {
	db := r.orm.Model(&models.Client{}).Where("id", clientId).Updates(data)

	return db.Error
}

func (r *ClientRepository) Delete(clientId uint, ctx context.Context) error {
	db := r.orm.Delete(&models.Client{}, clientId)
	return db.Error
}
