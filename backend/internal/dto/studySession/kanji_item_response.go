package studySessionDTO

import "github.com/google/uuid"

type KanjiItemResponse struct {
	ID        uuid.UUID `json:"id"`
	Character string    `json:"character"`
	Meaning   *string   `json:"meaning,omitempty"`
}