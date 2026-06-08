package models

import "github.com/google/uuid"

type Grammar struct {
	ID uuid.UUID `gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	UserID uuid.UUID `gorm:"type:uuid;not null"`
	User User `gorm:"foreignKey:UserID"`
	Pattern string `gorm:"size:255;not null"`
	Meaning *string
	Example *string
	Note *string
	Favourite bool `gorm:"default:false"`
	BaseModel
}