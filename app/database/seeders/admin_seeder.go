package seeders

import (
	"sso-auth/app/facades"
	"sso-auth/app/models"
)

type AdminSeeder struct {
}

func (s *AdminSeeder) Run() {
	var admin *models.Admin

	err := facades.Orm().Where("email = ?", "admin@admin.com").First(&admin)

	if err.Error == nil {
		return
	}

	adminUser := models.Admin{
		Email:    "admin@admin.com",
		Password: "password",
	}

	facades.Orm().Create(&adminUser)
}
