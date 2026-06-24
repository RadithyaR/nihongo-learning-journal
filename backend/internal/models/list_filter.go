package models

import "github.com/google/uuid"

type ListFilter struct {
	Search    string
	Favourite *bool
	TagID     *uuid.UUID
	SortBy    string // "newest", "oldest", "" (default)
}
