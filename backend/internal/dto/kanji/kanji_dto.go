package kanjiDTO

import "github.com/google/uuid"

type CreateKanjiRequest struct {
	Character string  `json:"character" validate:"required,max=10"`
	Meaning   *string `json:"meaning"`
	Onyomi    *string `json:"onyomi"`
	Kunyomi   *string `json:"kunyomi"`
	JLPTLevel *string `json:"jlpt_level"`
}

type UpdateKanjiRequest struct {
	Character string  `json:"character" validate:"required,max=10"`
	Meaning   *string `json:"meaning"`
	Onyomi    *string `json:"onyomi"`
	Kunyomi   *string `json:"kunyomi"`
	JLPTLevel *string `json:"jlpt_level"`
}

type KanjiResponse struct {
	ID         uuid.UUID `json:"id"`
	Character  string    `json:"character"`
	Meaning    *string   `json:"meaning"`
	Onyomi     *string   `json:"onyomi"`
	Kunyomi    *string   `json:"kunyomi"`
	JLPTLevel  *string   `json:"jlpt_level"`
	Favourite  bool      `json:"favourite"`
}