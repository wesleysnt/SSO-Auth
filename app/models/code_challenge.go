package models

import (
	"time"

	"gorm.io/gorm"
)

type CodeChallenge struct {
	ID         uint `gorm:"primaryKey"`
	Code       string
	UniqueCode *string `gorm:"uniqueIndex;default:uuid_generate_v4()"`
	Method     string
	ClientId   uint
	CreatedAt  time.Time
	UpdatedAt  time.Time
	DeletedAt  gorm.DeletedAt
}
