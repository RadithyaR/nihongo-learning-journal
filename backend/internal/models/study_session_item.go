package models

import "github.com/google/uuid"

type StudySessionItem struct {
	ID uuid.UUID `gorm:"type:uuid;primaryKey"`
	StudySessionID uuid.UUID `gorm:"type:uuid;not null;index"`
	ItemType string `gorm:"type:varchar(50);not null;index"`
	ItemID uuid.UUID `gorm:"type:uuid;not null;index"`
	StudySession StudySession `gorm:"foreignKey:StudySessionID"`
	BaseModel
}