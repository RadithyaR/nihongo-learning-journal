package models

import "github.com/google/uuid"

type Vocabulary struct {
	ID uuid.UUID `gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	UserID uuid.UUID `gorm:"type:uuid;not null"`
	User User `gorm:"foreignKey:UserID"`
	Word string `gorm:"size:255;not null"`
	Reading *string
	Source *string
	Note *string
	Status *string `gorm:"size:20;default:NEW"`
	Favourite bool `gorm:"default:false"`
	Meanings []VocabularyMeaning `gorm:"foreignKey:VocabularyID"`
	BaseModel
}