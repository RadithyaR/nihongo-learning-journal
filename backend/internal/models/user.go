package models

import "github.com/google/uuid"

type User struct {
	ID uuid.UUID `gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	Name string `gorm:"size:100;not null"`
	Email string `gorm:"size:225;uniqueIndex;not null"`
	PasswordHash string `gorm:"not null"` 
	GoogleID *string `gorm:"uniqueIndex"`
	AvatarURL *string
	BaseModel
}