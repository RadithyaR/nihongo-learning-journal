package models

import (
	"time"

	"github.com/google/uuid"
)

type Goal struct {
	ID uuid.UUID `gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	UserID uuid.UUID `gorm:"type:uuid;not null;index"`
	Title string `gorm:"size:255;not null"`
	Description *string `gorm:"type:text"`
	GoalType *string `gorm:"size:20"`
	TargetLevel *string `gorm:"size:10"`
	TargetCount *int
	TargetDate time.Time `gorm:"type:date;not null"`
	Status string `gorm:"size:20;default:IN_PROGRESS"`
	User User `gorm:"foreignKey:UserID"`
	BaseModel
}