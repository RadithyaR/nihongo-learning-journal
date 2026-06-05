package models

import (
	"time"

	"github.com/google/uuid"
)

type PasswordReset struct {
	ID uuid.UUID `gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	UserID uuid.UUID `gorm:"type:uuid;not null"`
	Token string `gorm:"size:255;uniqueIndex;not null"`
	ExpiresAt time.Time `gorm:"not null"`
	UsedAt *time.Time
	User User `gorm:"foreignKey:UserID"`
	CreatedAt time.Time
	UpdatedAt time.Time
}