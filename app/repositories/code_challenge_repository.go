package repositories

import (
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

func (r *CodeChallengeRepository) Create(codeChallenge *models.CodeChallenge) error {
	return r.orm.Create(&codeChallenge).Error
}

func (r *CodeChallengeRepository) GetChallenge(uniqueCode string, data *models.CodeChallenge) error {
	return r.orm.Where("unique_code = ?", uniqueCode).First(&data).Error
}
