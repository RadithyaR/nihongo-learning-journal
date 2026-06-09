package studySessionDTO

import "github.com/google/uuid"

type GrammarItemResponse struct {
	ID      uuid.UUID `json:"id"`
	Pattern string    `json:"pattern"`
	Meaning *string   `json:"meaning,omitempty"`
}