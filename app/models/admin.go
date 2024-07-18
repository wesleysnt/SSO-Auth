package models

import (
	"context"
	"sso-auth/app/utils"
	"time"

	"gorm.io/gorm"
)

type Admin struct {
	ID        uint   `gorm:"primarKey"`
	Email     string `gorm:"uniqueIndex"`
	Password  string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt
}

func (m *Admin) BeforeCreate(tx *gorm.DB) error {
	ctx := context.TODO()
	hashedPass, _ := utils.HashPassword(m.Password, ctx)

	m.Password = hashedPass

	return nil

}
