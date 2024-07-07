package repositories

import (
	"sso-auth/app/facades"
	"sso-auth/app/models"

	"gorm.io/gorm"
)

type OtpRequestRepository struct {
	orm *gorm.DB
}

func NewOtpRequestRepository() *OtpRequestRepository {
	return &OtpRequestRepository{orm: facades.Orm()}
}

func (r *OtpRequestRepository) GetByUniqueCode(uniqueCode string) (*models.OtpRequests, error) {
	var data models.OtpRequests
	err := r.orm.Where("unique_code = ?", uniqueCode).First(&data)
	return &data, err.Error
}

func (r *OtpRequestRepository) Store(data *models.OtpRequests) (*models.OtpRequests, error) {
	err := r.orm.Create(&data)
	return data, err.Error
}

func (r *OtpRequestRepository) Delete(uniqueCode string) error {
	res := r.orm.Where("unique_code = ?", uniqueCode).Delete(&models.OtpRequests{})
	return res.Error
}

func (r *OtpRequestRepository) DeleteByCustId(customerId uint) error {
	return r.orm.Where("customer_id = ?", customerId).Delete(&models.OtpRequests{}).Error
}
