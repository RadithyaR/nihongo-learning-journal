package errors

import "errors"

var (
	ErrGrammarAlreadyExists = errors.New("grammar already exists")
	ErrGrammarNotFound      = errors.New("grammar not found")
)