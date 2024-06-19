package models

import (
	"time"

	"gorm.io/gorm"
)

type AccessToken struct {
	ID         uint `gorm:"primaryKey"`
	Token      string
	UserId     uint
	ClientId   uint
	ExpiryTime time.Time
	Scope      string
	CreatedAt  time.Time
	UpdatedAt  time.Time
	DeletedAt  gorm.DeletedAt
}
