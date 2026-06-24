package grammar

import "github.com/google/uuid"

type CreateGrammarRequest struct {
	Pattern string  `json:"pattern" validate:"required"`
	Meaning *string `json:"meaning"`
	Example  *string `json:"example"`
	Note     *string `json:"note"`
	ImageURL *string `json:"image_url"`
}

type UpdateGrammarRequest struct {
	Pattern string  `json:"pattern" validate:"required"`
	Meaning *string `json:"meaning"`
	Example  *string `json:"example"`
	Note     *string `json:"note"`
	ImageURL *string `json:"image_url"`
}

type GrammarResponse struct {
	ID        uuid.UUID `json:"id"`
	Pattern   string    `json:"pattern"`
	Meaning   *string   `json:"meaning"`
	Example   *string   `json:"example"`
	Note      *string   `json:"note"`
	ImageURL  *string   `json:"image_url"`
	Favourite bool      `json:"favourite"`
}