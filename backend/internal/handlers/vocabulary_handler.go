package handlers

import (
	"net/http"

	vocabularyDTO "github.com/RadithyaR/nihongo-learning-journal/backend/internal/dto/vocabulary"
	serviceInterfaces "github.com/RadithyaR/nihongo-learning-journal/backend/internal/repositories/interfaces"
	customErrors "github.com/RadithyaR/nihongo-learning-journal/backend/pkg/errors"
	"github.com/RadithyaR/nihongo-learning-journal/backend/pkg/responses"
	appValidator "github.com/RadithyaR/nihongo-learning-journal/backend/pkg/validator"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type VocabularyHandler struct {
	vocabularyService serviceInterfaces.VocabularyService
}

func NewVocabularyHandler(
	vocabularyService serviceInterfaces.VocabularyService,
) *VocabularyHandler {
	return &VocabularyHandler{
		vocabularyService: vocabularyService,
	}
}

func (h *VocabularyHandler) CreateVocabulary(
	c *gin.Context,
) {

	var dto vocabularyDTO.CreateVocabularyRequest

	if err := c.ShouldBindJSON(&dto); err != nil {
		responses.Error(
			c,
			http.StatusBadRequest,
			"invalid request body",
		)
		return
	}

	if err := appValidator.Validate.Struct(dto); err != nil {
		responses.Error(
			c,
			http.StatusBadRequest,
			err.Error(),
		)
		return
	}

	userIDValue, exists := c.Get("user_id")

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

	vocabulary, err :=
		h.vocabularyService.CreateVocabulary(
			c.Request.Context(),
			userID,
			dto,
		)

	if err != nil {

		switch err {

		case customErrors.ErrVocabularyAlreadyExists:

			responses.Error(
				c,
				http.StatusConflict,
				err.Error(),
			)

		default:

			responses.Error(
				c,
				http.StatusInternalServerError,
				"internal server error",
			)
		}

		return
	}

	responses.Success(
		c,
		http.StatusCreated,
		"vocabulary created successfully",
		vocabulary,
	)
}

func (h *VocabularyHandler) GetVocabularyList(
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
	search := c.Query("search")

	vocabularies, err :=
		h.vocabularyService.GetVocabularyList(
			c.Request.Context(),
			userID,
			search,
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
		"vocabularies retrieved successfully",
		vocabularies,
	)
}

func (h *VocabularyHandler) GetVocabularyByID(
	c *gin.Context,
) {

	idParam := c.Param(
		"id",
	)

	vocabularyID, err := uuid.Parse(
		idParam,
	)

	if err != nil {
		responses.Error(
			c,
			http.StatusBadRequest,
			"invalid vocabulary id",
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

	vocabulary, err :=
		h.vocabularyService.GetVocabularyByID(
			c.Request.Context(),
			userID,
			vocabularyID,
		)

	if err != nil {

		switch err {

		case customErrors.ErrVocabularyNotFound:

			responses.Error(
				c,
				http.StatusNotFound,
				err.Error(),
			)

		default:

			responses.Error(
				c,
				http.StatusInternalServerError,
				"internal server error",
			)
		}

		return
	}

	responses.Success(
		c,
		http.StatusOK,
		"vocabulary retrieved successfully",
		vocabulary,
	)
}

func (h *VocabularyHandler) UpdateVocabulary(
	c *gin.Context,
) {

	idParam := c.Param(
		"id",
	)

	vocabularyID, err := uuid.Parse(
		idParam,
	)

	if err != nil {
		responses.Error(
			c,
			http.StatusBadRequest,
			"invalid vocabulary id",
		)
		return
	}

	var dto vocabularyDTO.UpdateVocabularyRequest

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

	vocabulary, err :=
		h.vocabularyService.UpdateVocabulary(
			c.Request.Context(),
			userID,
			vocabularyID,
			dto,
		)

	if err != nil {

		switch err {

		case customErrors.ErrVocabularyNotFound:

			responses.Error(
				c,
				http.StatusNotFound,
				err.Error(),
			)

		default:

			responses.Error(
				c,
				http.StatusInternalServerError,
				"internal server error",
			)
		}

		return
	}

	responses.Success(
		c,
		http.StatusOK,
		"vocabulary updated successfully",
		vocabulary,
	)
}

func (h *VocabularyHandler) DeleteVocabulary(
	c *gin.Context,
) {

	idParam := c.Param(
		"id",
	)

	vocabularyID, err := uuid.Parse(
		idParam,
	)

	if err != nil {

		responses.Error(
			c,
			http.StatusBadRequest,
			"invalid vocabulary id",
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

	err = h.vocabularyService.DeleteVocabulary(
		c.Request.Context(),
		userID,
		vocabularyID,
	)

	if err != nil {

		switch err {

		case customErrors.ErrVocabularyNotFound:

			responses.Error(
				c,
				http.StatusNotFound,
				err.Error(),
			)

		default:

			responses.Error(
				c,
				http.StatusInternalServerError,
				"internal server error",
			)
		}

		return
	}

	responses.Success(
		c,
		http.StatusOK,
		"vocabulary deleted successfully",
		nil,
	)
}

func (h *VocabularyHandler) ToggleFavourite(
	c *gin.Context,
) {

	idParam := c.Param(
		"id",
	)

	vocabularyID, err := uuid.Parse(
		idParam,
	)

	if err != nil {

		responses.Error(
			c,
			http.StatusBadRequest,
			"invalid vocabulary id",
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

	vocabulary, err := h.vocabularyService.ToggleFavourite(
		c.Request.Context(),
		userID,
		vocabularyID,
	)

	if err != nil {

		switch err {

		case customErrors.ErrVocabularyNotFound:

			responses.Error(
				c,
				http.StatusNotFound,
				err.Error(),
			)

		default:

			responses.Error(
				c,
				http.StatusInternalServerError,
				"internal server error",
			)
		}

		return
	}

	responses.Success(
		c,
		http.StatusOK,
		"vocabulary favourite updated successfully",
		vocabulary,
	)
}