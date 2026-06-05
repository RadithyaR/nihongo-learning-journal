package profile

import "github.com/google/uuid"

type ProfileResponse struct {
	ID         uuid.UUID `json:"id"`
	Name       string    `json:"name"`
	Email      string    `json:"email"`
	AvatarURL  *string   `json:"avatar_url"`
	IsVerified bool      `json:"is_verified"`
}

type UpdateProfileRequest struct {
	Name      string  `json:"name" validate:"required,min=3,max=100"`
	AvatarURL *string `json:"avatar_url"`
}