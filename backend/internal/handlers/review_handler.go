package handlers

import (
	"net/http"

	reviewDTO "github.com/RadithyaR/nihongo-learning-journal/backend/internal/dto/review"
	serviceInterfaces "github.com/RadithyaR/nihongo-learning-journal/backend/internal/repositories/interfaces"
	"github.com/RadithyaR/nihongo-learning-journal/backend/pkg/responses"
	appValidator "github.com/RadithyaR/nihongo-learning-journal/backend/pkg/validator"
	"github.com/gin-gonic/gin"
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

	userID, ok := GetUserID(c)

	if !ok {
		responses.Error(
			c,
			http.StatusUnauthorized,
			"user not found",
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

	userID, ok := GetUserID(c)

	if !ok {
		responses.Error(
			c,
			http.StatusUnauthorized,
			"user not found",
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
		nil,
	)
}