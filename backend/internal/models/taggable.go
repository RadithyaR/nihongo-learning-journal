package models

import "github.com/google/uuid"

type Taggable struct {
	ID uuid.UUID `gorm:"type:uuid;default:gen_random_uuid();primaryKey"`

	TagID uuid.UUID `gorm:"type:uuid;not null;index"`

	ItemType string `gorm:"size:20;not null"`

	ItemID uuid.UUID `gorm:"type:uuid;not null"`

	Tag Tag `gorm:"foreignKey:TagID"`

	BaseModel
}