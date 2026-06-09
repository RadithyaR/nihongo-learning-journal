package models

import (
	"time"

	"github.com/google/uuid"
)

type StudySession struct {
	ID uuid.UUID `gorm:"type:uuid;primaryKey"`
	UserID uuid.UUID `gorm:"type:uuid;not null;index"`
	SessionDate time.Time `gorm:"type:date;not null;index"`
	Notes string `gorm:"type:text"`
	Reflection string `gorm:"type:text"`
	User User `gorm:"foreignKey:UserID"`
	BaseModel
}