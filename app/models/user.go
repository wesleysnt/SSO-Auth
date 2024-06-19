package models

import (
	"sso-auth/app/utils"
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID        uint   `gorm:"primarKey"`
	Username  string `gorm:"uniqueIndex"`
	Email     string `gorm:"uniqueIndex"`
	Password  string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt
}

func (m *User) BeforeCreate(tx *gorm.DB) error {
	hashedPass, _ := utils.HashPassword(m.Password)

	m.Password = hashedPass

	return nil

}
