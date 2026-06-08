package reviewdto

import "github.com/google/uuid"

type NextKanjiReviewResponse struct {
	ID        uuid.UUID `json:"id"`
	Character string    `json:"character"`
	Meaning   *string   `json:"meaning"`
	Onyomi    *string   `json:"onyomi"`
	Kunyomi   *string   `json:"kunyomi"`
}