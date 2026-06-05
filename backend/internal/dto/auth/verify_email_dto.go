package auth

type VerifyEmailDTO struct {
	Token string `json:"token" validate:"required"`
}