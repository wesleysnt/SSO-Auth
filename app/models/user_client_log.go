package models

import (
	"time"

	"gorm.io/gorm"
)

type UserClientLog struct {
	ID        uint `gorm:"primaryKey"`
	UserId    uint
	ClientId  uint
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt
}
