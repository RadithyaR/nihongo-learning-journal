package models

import (
	"time"

	"github.com/google/uuid"
)

type SRSRecord struct {
	ID uuid.UUID `gorm:"type:uuid;default:gen_random_uuid();primaryKey"`

	UserID uuid.UUID `gorm:"type:uuid;not null;index"`
	User User `gorm:"foreignKey:UserID"`

	ItemType string `gorm:"size:20;not null;index"`
	ItemID uuid.UUID `gorm:"type:uuid;not null;index"`

	EaseFactor float64 `gorm:"default:2.5"`

	IntervalDays int `gorm:"default:0"`

	ReviewCount int `gorm:"default:0"`

	LastReviewedAt *time.Time

	NextReviewAt time.Time `gorm:"not null;index"`

	BaseModel
}