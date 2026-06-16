package services

import (
	"context"
	"time"

	"github.com/RadithyaR/nihongo-learning-journal/backend/internal/constants"
	"github.com/RadithyaR/nihongo-learning-journal/backend/internal/models"
	repositoryInterfaces "github.com/RadithyaR/nihongo-learning-journal/backend/internal/repositories/interfaces"
	serviceInterfaces "github.com/RadithyaR/nihongo-learning-journal/backend/internal/services/interfaces"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type srsService struct {
	srsRepository repositoryInterfaces.SRSRepository
}

func NewSRSService(
	srsRepository repositoryInterfaces.SRSRepository,
) serviceInterfaces.SRSService {

	return &srsService{
		srsRepository: srsRepository,
	}
}

func (s *srsService) calculateInterval(
	rating string,
	reviewCount int,
) int {

	switch rating {

	case constants.RatingAgain:
		return 1

	case constants.RatingHard:

		if reviewCount == 0 {
			return 3
		}

		return reviewCount * 3

	case constants.RatingGood:

		if reviewCount == 0 {
			return 7
		}

		return reviewCount * 7

	case constants.RatingEasy:

		if reviewCount == 0 {
			return 14
		}

		return reviewCount * 14
	}

	return 1
}

func (s *srsService) UpdateReview(
	ctx context.Context,
	userID uuid.UUID,
	itemType string,
	itemID uuid.UUID,
	rating string,
) error {

	record, err := s.srsRepository.FindByItem(
		ctx,
		userID,
		itemType,
		itemID,
	)

	if err != nil {

		if err == gorm.ErrRecordNotFound {

			now := time.Now()

			interval := s.calculateInterval(
				rating,
				0,
			)

			newRecord := models.SRSRecord{
				UserID: userID,

				ItemType: itemType,
				ItemID: itemID,

				EaseFactor: 2.5,

				IntervalDays: interval,

				ReviewCount: 1,

				LastReviewedAt: &now,

				NextReviewAt: now.AddDate(
					0,
					0,
					interval,
				),
			}

			return s.srsRepository.Create(
				ctx,
				&newRecord,
			)
		}

		return err
	}

	now := time.Now()

	record.ReviewCount++

	record.IntervalDays = s.calculateInterval(
		rating,
		record.ReviewCount,
	)

	record.LastReviewedAt = &now

	record.NextReviewAt = now.AddDate(
		0,
		0,
		record.IntervalDays,
	)

	return s.srsRepository.Update(
		ctx,
		record,
	)
}

func (s *srsService) GetDueCount(
	ctx context.Context,
	userID uuid.UUID,
) (int64, error) {

	return s.srsRepository.CountDueToday(
		ctx,
		userID,
	)
}

func (s *srsService) GetOverdueCount(
	ctx context.Context,
	userID uuid.UUID,
) (int64, error) {

	return s.srsRepository.CountOverdue(
		ctx,
		userID,
	)
}

func (s *srsService) GetNextDueItem(
	ctx context.Context,
	userID uuid.UUID,
	itemType string,
) (*models.SRSRecord, error) {

	return s.srsRepository.GetNextDueItem(
		ctx,
		userID,
		itemType,
	)
}
