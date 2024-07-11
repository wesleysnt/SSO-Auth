package models

import (
	"time"

	"gorm.io/gorm"
)

type ForgotPasswordToken struct {
	ID         uint `gorm:"primaryKey"`
	UserId     uint
	Token      string
	ExpiryTime time.Time
	IsUsed     bool
	CreatedAt  time.Time
	UpdatedAt  time.Time
	DeletedAt  gorm.DeletedAt
}
