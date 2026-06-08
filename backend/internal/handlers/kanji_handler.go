package handlers

import (
	"net/http"

	kanjiDTO "github.com/RadithyaR/nihongo-learning-journal/backend/internal/dto/kanji"
	serviceInterfaces "github.com/RadithyaR/nihongo-learning-journal/backend/internal/repositories/interfaces"
	customErrors "github.com/RadithyaR/nihongo-learning-journal/backend/pkg/errors"
	"github.com/RadithyaR/nihongo-learning-journal/backend/pkg/responses"
	appValidator "github.com/RadithyaR/nihongo-learning-journal/backend/pkg/validator"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type KanjiHandler struct {
	kanjiService serviceInterfaces.KanjiService
}

func NewKanjiHandler(
	kanjiService serviceInterfaces.KanjiService,
) *KanjiHandler {

	return &KanjiHandler{
		kanjiService: kanjiService,
	}
}


func (h *KanjiHandler) CreateKanji(
	c *gin.Context,
) {

	var dto kanjiDTO.CreateKanjiRequest

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

	kanji, err := h.kanjiService.CreateKanji(
		c.Request.Context(),
		userID,
		dto,
	)

	if err != nil {

		switch err {

		case customErrors.ErrKanjiAlreadyExists:

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
		"kanji created successfully",
		kanji,
	)
}

func (h *KanjiHandler) GetKanjiList(
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

	search := c.Query("search")

	kanjis, err := h.kanjiService.GetKanjiList(
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
		"kanji list",
		kanjis,
	)
}

func (h *KanjiHandler) GetKanjiByID(
	c *gin.Context,
) {

	idParam := c.Param("id")

	id, err := uuid.Parse(
		idParam,
	)

	if err != nil {

		responses.Error(
			c,
			http.StatusBadRequest,
			"invalid id",
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

	kanji, err := h.kanjiService.GetKanjiByID(
		c.Request.Context(),
		userID,
		id,
	)

	if err != nil {

		responses.Error(
			c,
			http.StatusNotFound,
			err.Error(),
		)

		return
	}

	responses.Success(
		c,
		http.StatusOK,
		"kanji detail",
		kanji,
	)
}

func (h *KanjiHandler) UpdateKanji(
	c *gin.Context,
) {

	var dto kanjiDTO.UpdateKanjiRequest

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

	id, err := uuid.Parse(
		c.Param("id"),
	)

	if err != nil {

		responses.Error(
			c,
			http.StatusBadRequest,
			"invalid id",
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

	kanji, err := h.kanjiService.UpdateKanji(
		c.Request.Context(),
		userID,
		id,
		dto,
	)

	if err != nil {

		responses.Error(
			c,
			http.StatusBadRequest,
			err.Error(),
		)

		return
	}

	responses.Success(
		c,
		http.StatusOK,
		"kanji updated successfully",
		kanji,
	)
}

func (h *KanjiHandler) DeleteKanji(
	c *gin.Context,
) {

	id, err := uuid.Parse(
		c.Param("id"),
	)

	if err != nil {

		responses.Error(
			c,
			http.StatusBadRequest,
			"invalid id",
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

	err = h.kanjiService.DeleteKanji(
		c.Request.Context(),
		userID,
		id,
	)

	if err != nil {

		responses.Error(
			c,
			http.StatusNotFound,
			err.Error(),
		)

		return
	}

	responses.Success(
		c,
		http.StatusOK,
		"kanji deleted successfully",
		nil,
	)
}

func (h *KanjiHandler) ToggleFavourite(
	c *gin.Context,
) {

	id, err := uuid.Parse(
		c.Param("id"),
	)

	if err != nil {

		responses.Error(
			c,
			http.StatusBadRequest,
			"invalid id",
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

	kanji, err := h.kanjiService.ToggleFavourite(
		c.Request.Context(),
		userID,
		id,
	)

	if err != nil {

		responses.Error(
			c,
			http.StatusNotFound,
			err.Error(),
		)

		return
	}

	responses.Success(
		c,
		http.StatusOK,
		"kanji favourite updated",
		kanji,
	)
}

