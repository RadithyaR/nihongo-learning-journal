package studySessionDTO

import "github.com/google/uuid"

type VocabularyItemResponse struct {
	ID      uuid.UUID `json:"id"`
	Word    string    `json:"word"`
	Reading *string   `json:"reading,omitempty"`
}