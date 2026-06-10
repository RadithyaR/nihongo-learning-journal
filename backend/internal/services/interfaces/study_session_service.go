package interfaces

import (
	"context"

	studySessionDTO "github.com/RadithyaR/nihongo-learning-journal/backend/internal/dto/studySession"
	"github.com/google/uuid"
)

type StudySessionService interface {
	GetTodaySession(
		ctx context.Context,
		userID uuid.UUID,
	) (*studySessionDTO.TodaySessionResponse, error)

	UpdateNotes(
		ctx context.Context,
		userID uuid.UUID,
		notes string,
	) error

	UpdateReflection(
		ctx context.Context,
		userID uuid.UUID,
		reflection string,
	) error

	AddVocabulary(
		ctx context.Context,
		userID uuid.UUID,
		vocabularyID uuid.UUID,
	) error

	AddKanji(
		ctx context.Context,
		userID uuid.UUID,
		kanjiID uuid.UUID,
	) error

	AddGrammar(
		ctx context.Context,
		userID uuid.UUID,
		grammarID uuid.UUID,
	) error
}