package models

import (
	"time"

	"gorm.io/gorm"
)

type RefreshToken struct {
	ID         uint `gorm:"primaryKey"`
	Token      string
	UserId     uint
	ClientId   uint
	ExpiryTime time.Time
	CreatedAt  time.Time
	UpdatedAt  time.Time
	DeletedAt  gorm.DeletedAt
}
