package models

import "github.com/google/uuid"

type Tag struct {
	ID uuid.UUID `gorm:"type:uuid;default:gen_random_uuid();primaryKey"`

	UserID uuid.UUID `gorm:"type:uuid;not null;index"`

	Name string `gorm:"size:100;not null"`

	Color *string `gorm:"size:20"`

	User User `gorm:"foreignKey:UserID"`

	BaseModel
}