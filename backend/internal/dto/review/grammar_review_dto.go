package review

import "github.com/google/uuid"

type NextGrammarReviewResponse struct {
	ID      uuid.UUID `json:"id"`
	Pattern string    `json:"pattern"`
	Meaning *string   `json:"meaning"`
}

type SubmitGrammarReviewRequest struct {
	GrammarID uuid.UUID `json:"grammar_id" validate:"required"`
	Rating    string    `json:"rating" validate:"required,oneof=AGAIN HARD GOOD EASY"`
}