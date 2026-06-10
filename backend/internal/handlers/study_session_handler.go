package handlers

import (
	"net/http"

	studySessionDTO "github.com/RadithyaR/nihongo-learning-journal/backend/internal/dto/studySession"
	serviceInterfaces "github.com/RadithyaR/nihongo-learning-journal/backend/internal/services/interfaces"
	"github.com/RadithyaR/nihongo-learning-journal/backend/pkg/responses"
	"github.com/gin-gonic/gin"
)

type StudySessionHandler struct {
	studySessionService serviceInterfaces.StudySessionService
}

func NewStudySessionHandler(
	studySessionService serviceInterfaces.StudySessionService,
) *StudySessionHandler {

	return &StudySessionHandler{
		studySessionService: studySessionService,
	}
}

func (h *StudySessionHandler) GetTodaySession(
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

	session, err :=
		h.studySessionService.GetTodaySession(
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
		"study session retrieved successfully",
		session,
	)
}

func (h *StudySessionHandler) UpdateNotes(
	c *gin.Context,
) {

	var dto studySessionDTO.UpdateNotesRequest

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

	userID, ok := GetUserID(c)

	if !ok {

		responses.Error(
			c,
			http.StatusUnauthorized,
			"user not found",
		)

		return
	}

	err := h.studySessionService.UpdateNotes(
		c.Request.Context(),
		userID,
		dto.Notes,
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
		"notes updated successfully",
		nil,
	)
}

func (h *StudySessionHandler) UpdateReflection(
	c *gin.Context,
) {

	var dto studySessionDTO.UpdateReflectionRequest

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

	userID, ok := GetUserID(c)

	if !ok {

		responses.Error(
			c,
			http.StatusUnauthorized,
			"user not found",
		)

		return
	}

	err := h.studySessionService.UpdateReflection(
		c.Request.Context(),
		userID,
		dto.Reflection,
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
		"reflection updated successfully",
		nil,
	)
}