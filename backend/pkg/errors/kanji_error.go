package errors

import "errors"

var (
	ErrKanjiNotFound = errors.New(
		"kanji not found",
	)

	ErrKanjiAlreadyExists = errors.New(
		"kanji already exists",
	)
)