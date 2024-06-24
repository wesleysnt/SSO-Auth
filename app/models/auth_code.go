package models

import (
	"time"

	"gorm.io/gorm"
)

type AuthCode struct {
	ID         uint `gorm:"primaryKey"`
	Code       *string
	ExpiryTime time.Time
	ClientId   uint
	UserId     uint
	CreatedAt  time.Time
	UpdatedAt  time.Time
	DeletedAt  gorm.DeletedAt
}
