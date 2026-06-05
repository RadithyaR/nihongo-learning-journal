package auth

type ResetPasswordDTO struct {
	Token string `json:"token" validate:"required"`

	NewPassword string `json:"new_password" validate:"required,min=8"`
}