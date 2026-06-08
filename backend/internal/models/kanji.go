package models

import "github.com/google/uuid"

type Kanji struct {
	ID uuid.UUID `gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	UserID uuid.UUID `gorm:"type:uuid;not null"`
	User User `gorm:"foreignKey:UserID"`
	Character string `gorm:"size:10;not null"`
	Meaning *string
	Onyomi *string
	Kunyomi *string
	JLPTLevel *string `gorm:"size:10"`
	Favourite bool `gorm:"default:false"`
	BaseModel
}