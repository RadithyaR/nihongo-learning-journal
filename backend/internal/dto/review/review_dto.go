package reviewdto

import "github.com/google/uuid"

type SubmitReviewRequest struct {
	VocabularyID uuid.UUID `json:"vocabulary_id" validate:"required"`
	Rating       string    `json:"rating" validate:"required,oneof=AGAIN HARD GOOD EASY"`
}

type NextReviewResponse struct {
	ID       uuid.UUID `json:"id"`
	Word     string    `json:"word"`
	Reading  *string   `json:"reading"`
}