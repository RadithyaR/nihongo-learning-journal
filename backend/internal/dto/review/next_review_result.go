package review

type NextReviewResult struct {
	HasNext bool                  `json:"has_next"`
	Review  *NextReviewResponse   `json:"review,omitempty"`
}

type NextKanjiReviewResult struct {
	HasNext bool                      `json:"has_next"`
	Review  *NextKanjiReviewResponse  `json:"review,omitempty"`
}

type NextGrammarReviewResult struct {
	HasNext bool                       `json:"has_next"`
	Review  *NextGrammarReviewResponse `json:"review,omitempty"`
}
