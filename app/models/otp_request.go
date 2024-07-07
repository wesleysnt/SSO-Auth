package models

import (
	"time"

	"gorm.io/gorm"
)

type OtpRequests struct {
	ID         uint `gorm:"primaryKey"`
	OtpCode    string
	UniqueCode string
	IsUsed     bool
	UserId     uint
	ExpiredAt  time.Time
	CreatedAt  time.Time
	UpdatedAt  time.Time
	DeletedAt  gorm.DeletedAt
}
