package vocabulary

type CreateVocabularyDTO struct {
	Word    string   `json:"word" validate:"required"`
	Reading string   `json:"reading"`
	Meaning []string `json:"meanings" validate:"required,min=1"`
	Source  string   `json:"source"`
	Note    string   `json:"note"`
	Status  string   `json:"status"`
}