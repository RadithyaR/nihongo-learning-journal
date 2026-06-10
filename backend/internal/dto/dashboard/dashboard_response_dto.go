package dashboard

import (
	"time"

	"github.com/google/uuid"
)

type RecentSessionResponse struct {
	ID          uuid.UUID    `json:"uuid;default:gen_random_uuid();primaryKey"`
	SessionDate time.Time `json:"sessionDate"`
	Notes       string    `json:"notes"`
}

type DashboardResponse struct {
	TotalVocabulary int `json:"totalVocabulary"`
	TotalKanji      int `json:"totalKanji"`
	TotalGrammar    int `json:"totalGrammar"`

	ReviewCountToday int `json:"reviewCountToday"`

	StudyStreak int `json:"studyStreak"`

	ActiveGoals    int `json:"activeGoals"`
	CompletedGoals int `json:"completedGoals"`

	RecentSessions []RecentSessionResponse `json:"recentSessions"`
}