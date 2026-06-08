package models

import (
	"time"

	"github.com/google/uuid"
)

type ReviewLog struct {
	ID uuid.UUID `gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	UserID uuid.UUID `gorm:"type:uuid;not null"`
	User User `gorm:"foreignKey:UserID"`
	VocabularyID uuid.UUID `gorm:"type:uuid;not null"`
	Vocabulary Vocabulary `gorm:"foreignKey:VocabularyID"`
	Rating string `gorm:"size:20;not null"`
	ReviewedAt time.Time `gorm:"not null"`
}