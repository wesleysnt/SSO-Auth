package repositories

import (
	"context"
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

func (r *OtpRequestRepository) GetByUniqueCode(uniqueCode string, ctx context.Context) (*models.OtpRequests, error) {
	var data models.OtpRequests
	err := r.orm.WithContext(ctx).Where("unique_code = ?", uniqueCode).First(&data)
	return &data, err.Error
}

func (r *OtpRequestRepository) Store(data *models.OtpRequests, ctx context.Context) (*models.OtpRequests, error) {
	err := r.orm.WithContext(ctx).Create(&data)
	return data, err.Error
}

func (r *OtpRequestRepository) Delete(uniqueCode string, ctx context.Context) error {
	res := r.orm.WithContext(ctx).Where("unique_code = ?", uniqueCode).Delete(&models.OtpRequests{})
	return res.Error
}

func (r *OtpRequestRepository) DeleteByCustId(customerId uint, ctx context.Context) error {
	return r.orm.WithContext(ctx).Where("customer_id = ?", customerId).Delete(&models.OtpRequests{}).Error
}
