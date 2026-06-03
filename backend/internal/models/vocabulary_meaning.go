package models

import (
	"time"

	"github.com/google/uuid"
)

type VocabularyMeaning struct {
	ID uuid.UUID `gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	VocabularyID uuid.UUID `gorm:"type:uuid;not null"`
	Meaning string `gorm:"not null"`
	OrderNumber int `gorm:"not null"`
	CreatedAt time.Time
	UpdatedAt time.Time
}