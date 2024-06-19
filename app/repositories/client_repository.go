package repositories

import (
	"sso-auth/app/facades"
	"sso-auth/app/models"

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

func (r *ClientRepository) Create(data *models.Client) error {
	res := r.orm.Create(&data)
	return res.Error
}
func (r *ClientRepository) GetById(data *models.Client, id uint) error {
	res := r.orm.First(&data, id)
	return res.Error
}

func (r *ClientRepository) CheckClientId(clientId string) error {
	var data *models.Client
	res := r.orm.Where("client_id = ?", clientId).First(&data)
	return res.Error
}
