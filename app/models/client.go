package models

import (
	"time"

	"gorm.io/gorm"
)

type Client struct {
	ID          uint    `gorm:"primaryKey"`
	ClientId    *string `gorm:"uniqueIndex;default:uuid_generate_v4()"`
	Name        string
	Secret      string
	RedirectUri string
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   gorm.DeletedAt
}
