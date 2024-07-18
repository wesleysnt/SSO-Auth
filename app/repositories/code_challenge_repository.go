package repositories

import (
	"context"
	"sso-auth/app/facades"
	"sso-auth/app/models"

	"gorm.io/gorm"
)

type CodeChallengeRepository struct {
	orm *gorm.DB
}

func NewCodeChallengeRepository() *CodeChallengeRepository {
	return &CodeChallengeRepository{orm: facades.Orm()}
}

func (r *CodeChallengeRepository) Create(codeChallenge *models.CodeChallenge, ctx context.Context) error {
	return r.orm.WithContext(ctx).Create(&codeChallenge).Error
}

func (r *CodeChallengeRepository) GetChallenge(uniqueCode string, data *models.CodeChallenge, ctx context.Context) error {
	return r.orm.WithContext(ctx).Where("unique_code = ?", uniqueCode).First(&data).Error
}
