package studySessionDTO

type TodaySessionResponse struct {
	Date string `json:"date"`

	Notes string `json:"notes"`

	Reflection string `json:"reflection"`

	Vocabularies []VocabularyItemResponse `json:"vocabularies"`

	Kanjis []KanjiItemResponse `json:"kanjis"`

	Grammars []GrammarItemResponse `json:"grammars"`
}