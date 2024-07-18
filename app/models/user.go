package models

import (
	"context"
	"sso-auth/app/utils"
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID        uint   `gorm:"primarKey"`
	Email     string `gorm:"uniqueIndex"`
	Password  string
	Name      string
	Phone     string
	IsActive  bool
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt
}

func (m *User) BeforeCreate(tx *gorm.DB) error {
	ctx := context.TODO()
	hashedPass, _ := utils.HashPassword(m.Password, ctx)

	m.Password = hashedPass

	return nil

}

func (m *User) BeforeUpdate(tx *gorm.DB) error {
	ctx := context.TODO()
	hashedPass, _ := utils.HashPassword(m.Password, ctx)

	tx.Statement.SetColumn("password", hashedPass)
	return nil

}
