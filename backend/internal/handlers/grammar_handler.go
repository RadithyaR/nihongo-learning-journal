package handlers

import (
	"net/http"

	grammarDTO "github.com/RadithyaR/nihongo-learning-journal/backend/internal/dto/grammar"
	"github.com/RadithyaR/nihongo-learning-journal/backend/internal/models"
	serviceInterfaces "github.com/RadithyaR/nihongo-learning-journal/backend/internal/services/interfaces"
	customErrors "github.com/RadithyaR/nihongo-learning-journal/backend/pkg/errors"
	"github.com/RadithyaR/nihongo-learning-journal/backend/pkg/responses"
	appValidator "github.com/RadithyaR/nihongo-learning-journal/backend/pkg/validator"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type GrammarHandler struct {
	grammarService serviceInterfaces.GrammarService
}

func NewGrammarHandler(
	grammarService serviceInterfaces.GrammarService,
) *GrammarHandler {

	return &GrammarHandler{
		grammarService: grammarService,
	}
}

func (h *GrammarHandler) CreateGrammar(
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

	var dto grammarDTO.CreateGrammarRequest

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

	grammar, err := h.grammarService.CreateGrammar(
		c.Request.Context(),
		userID,
		dto,
	)

	if err != nil {

		switch err {

		case customErrors.ErrGrammarAlreadyExists:
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
		"grammar created successfully",
		grammar,
	)
}

func (h *GrammarHandler) GetGrammarList(
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

	filter := models.ListFilter{
		Search: c.Query("search"),
		SortBy: c.Query("sort"),
	}

	if favStr := c.Query("favourite"); favStr != "" {
		fav := favStr == "true"
		filter.Favourite = &fav
	}

	if tagIDStr := c.Query("tagId"); tagIDStr != "" {
		if tagID, err := uuid.Parse(tagIDStr); err == nil {
			filter.TagID = &tagID
		}
	}

	grammars, err := h.grammarService.GetGrammarList(
		c.Request.Context(),
		userID,
		filter,
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
		"grammar list",
		grammars,
	)
}

func (h *GrammarHandler) GetGrammarByID(
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

	id, err := uuid.Parse(
		c.Param("id"),
	)

	if err != nil {
		responses.Error(
			c,
			http.StatusBadRequest,
			"invalid grammar id",
		)
		return
	}

	grammar, err := h.grammarService.GetGrammarByID(
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
		"grammar found",
		grammar,
	)
}

func (h *GrammarHandler) UpdateGrammar(
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

	id, err := uuid.Parse(
		c.Param("id"),
	)

	if err != nil {
		responses.Error(
			c,
			http.StatusBadRequest,
			"invalid grammar id",
		)
		return
	}

	var dto grammarDTO.UpdateGrammarRequest

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

	grammar, err := h.grammarService.UpdateGrammar(
		c.Request.Context(),
		userID,
		id,
		dto,
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
		"grammar updated successfully",
		grammar,
	)
}

func (h *GrammarHandler) DeleteGrammar(
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

	id, err := uuid.Parse(
		c.Param("id"),
	)

	if err != nil {
		responses.Error(
			c,
			http.StatusBadRequest,
			"invalid grammar id",
		)
		return
	}

	err = h.grammarService.DeleteGrammar(
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
		"grammar deleted successfully",
		nil,
	)
}
func (h *GrammarHandler) ToggleFavourite(
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

	id, err := uuid.Parse(
		c.Param("id"),
	)

	if err != nil {
		responses.Error(
			c,
			http.StatusBadRequest,
			"invalid grammar id",
		)
		return
	}

	grammar, err := h.grammarService.ToggleFavourite(
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
		"grammar favourite updated",
		grammar,
	)
}