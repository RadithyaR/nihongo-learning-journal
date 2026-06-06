package errors

import "errors"

var (
	ErrVocabularyAlreadyExists = errors.New("vocabulary already exists")
	ErrVocabularyNotFound      = errors.New("vocabulary not found")
)