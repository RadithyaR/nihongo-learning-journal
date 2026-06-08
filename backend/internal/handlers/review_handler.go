package handlers

import (
	"net/http"

	reviewDTO "github.com/RadithyaR/nihongo-learning-journal/backend/internal/dto/review"
	serviceInterfaces "github.com/RadithyaR/nihongo-learning-journal/backend/internal/repositories/interfaces"
	"github.com/RadithyaR/nihongo-learning-journal/backend/pkg/responses"
	appValidator "github.com/RadithyaR/nihongo-learning-journal/backend/pkg/validator"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type ReviewHandler struct {
	reviewService serviceInterfaces.ReviewService
}

func NewReviewHandler(
	reviewService serviceInterfaces.ReviewService,
) *ReviewHandler {

	return &ReviewHandler{
		reviewService: reviewService,
	}
}

func (h *ReviewHandler) GetNextReview(
	c *gin.Context,
) {

	userIDValue, exists := c.Get(
		"user_id",
	)

	if !exists {

		responses.Error(
			c,
			http.StatusUnauthorized,
			"user not found",
		)

		return
	}

	userID, ok := userIDValue.(uuid.UUID)

	if !ok {

		responses.Error(
			c,
			http.StatusUnauthorized,
			"invalid user",
		)

		return
	}

	review, err := h.reviewService.GetNextReview(
		c.Request.Context(),
		userID,
	)

	if err != nil {

		responses.Error(
			c,
			http.StatusInternalServerError,
			"internal server error",
		)

		return
	}

	responses.Success(
		c,
		http.StatusOK,
		"next review retrieved successfully",
		review,
	)
}

func (h *ReviewHandler) SubmitReview(
	c *gin.Context,
) {

	var dto reviewDTO.SubmitReviewRequest

	if err := c.ShouldBindJSON(
		&dto,
	); err != nil {

		responses.Error(
			c,
			http.StatusBadRequest,
			"invalid request body",
		)

		return
	}

	if err := appValidator.Validate.Struct(
		dto,
	); err != nil {

		responses.Error(
			c,
			http.StatusBadRequest,
			err.Error(),
		)

		return
	}

	userIDValue, exists := c.Get(
		"user_id",
	)

	if !exists {

		responses.Error(
			c,
			http.StatusUnauthorized,
			"user not found",
		)

		return
	}

	userID, ok := userIDValue.(uuid.UUID)

	if !ok {

		responses.Error(
			c,
			http.StatusUnauthorized,
			"invalid user",
		)

		return
	}

	err := h.reviewService.SubmitReview(
		c.Request.Context(),
		userID,
		dto,
	)

	if err != nil {

		responses.Error(
			c,
			http.StatusInternalServerError,
			"internal server error",
		)

		return
	}

	responses.Success(
		c,
		http.StatusCreated,
		"review submitted successfully",
		gin.H{
		"vocabulary_id": dto.VocabularyID,
		"rating": dto.Rating,
		},
	)
}