package handlers

import (
	"net/http"

	profileDTO "github.com/RadithyaR/nihongo-learning-journal/backend/internal/dto/profile"
	serviceInterfaces "github.com/RadithyaR/nihongo-learning-journal/backend/internal/services/interfaces"
	"github.com/RadithyaR/nihongo-learning-journal/backend/pkg/responses"
	appValidator "github.com/RadithyaR/nihongo-learning-journal/backend/pkg/validator"
	"github.com/gin-gonic/gin"
)

type ProfileHandler struct {
	profileService serviceInterfaces.ProfileService
}

func NewProfileHandler(
	profileService serviceInterfaces.ProfileService,
) *ProfileHandler {
	return &ProfileHandler{
		profileService: profileService,
	}
}

func (h *ProfileHandler) GetProfile(
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

	profile, err := h.profileService.GetProfile(
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
		"profile retrieved successfully",
		profile,
	)
}

func (h *ProfileHandler) UpdateProfile(
	c *gin.Context,
) {

	var dto profileDTO.UpdateProfileRequest

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

	profile, err := h.profileService.UpdateProfile(
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
		http.StatusOK,
		"profile updated successfully",
		profile,
	)
}