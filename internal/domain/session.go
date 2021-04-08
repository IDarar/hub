package domain

import (
	"time"

	_ "gorm.io/gorm"
)

type Session struct {
	UserID       int       `gorm:"primaryKey"`
	RefreshToken string    `json:"refreshToken" bson:"refreshToken"`
	ExpiresAt    time.Time `json:"expiresAt" bson:"expiresAt"`
}
