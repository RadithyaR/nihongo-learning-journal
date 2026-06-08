package vocabulary

import "github.com/google/uuid"

type CreateVocabularyRequest struct {
	Word     string   `json:"word" validate:"required,max=255"`
	Reading  *string  `json:"reading"`
	Source   *string  `json:"source"`
	Note     *string  `json:"note"`
	Meanings []string `json:"meanings" validate:"required,min=1"`
}

type UpdateVocabularyRequest struct {
	Word     string   `json:"word" validate:"required,max=255"`
	Reading  *string  `json:"reading"`
	Source   *string  `json:"source"`
	Note     *string  `json:"note"`
	Meanings []string `json:"meanings" validate:"required,min=1"`
}

type VocabularyMeaningResponse struct {
	ID          uuid.UUID `json:"id"`
	Meaning     string    `json:"meaning"`
	OrderNumber int       `json:"order_number"`
}

type VocabularyResponse struct {
	ID        uuid.UUID                   `json:"id"`
	Word      string                      `json:"word"`
	Reading   *string                     `json:"reading"`
	Source    *string                     `json:"source"`
	Note      *string                     `json:"note"`
	Status    *string                     `json:"status"`
	Favourite bool                        `json:"favourite"`
	Meanings  []VocabularyMeaningResponse `json:"meanings"`
}