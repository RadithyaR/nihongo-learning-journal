package analytics

type DailyStat struct {
	Date  string `json:"date"`
	Count int    `json:"count"`
}

type AnalyticsResponse struct {
	ReviewsPerDay    []DailyStat `json:"reviews_per_day"`
	VocabularyGrowth []DailyStat `json:"vocabulary_growth"`
	StudySessions    []DailyStat `json:"study_sessions"`
	CompletionRate   float64     `json:"completion_rate"`
}