package models

import (
	"time"

	"github.com/google/uuid"
)

type UserSession struct {
	ID uuid.UUID `gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	TokenID uuid.UUID `gorm:"type:uuid;uniqueIndex;not null"`
	UserID uuid.UUID `gorm:"type:uuid;not null"`
	User User `gorm:"foreignKey:UserID"`
	RefreshTokenHash string `gorm:"not null"`
	DeviceName *string
	IPAddress *string
	UserAgent *string
	ExpiresAt time.Time
	LastUsedAt *time.Time
	CreatedAt time.Time
	UpdatedAt time.Time
}