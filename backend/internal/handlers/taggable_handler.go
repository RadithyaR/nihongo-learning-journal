package handlers

import (
	"net/http"

	tagDTO "github.com/RadithyaR/nihongo-learning-journal/backend/internal/dto/tag"
	serviceInterfaces "github.com/RadithyaR/nihongo-learning-journal/backend/internal/services/interfaces"
	"github.com/RadithyaR/nihongo-learning-journal/backend/pkg/responses"
	appValidator "github.com/RadithyaR/nihongo-learning-journal/backend/pkg/validator"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type TaggableHandler struct {
	taggableService serviceInterfaces.TaggableService
}

func NewTaggableHandler(
	taggableService serviceInterfaces.TaggableService,
) *TaggableHandler {

	return &TaggableHandler{
		taggableService: taggableService,
	}
}

func (h *TaggableHandler) AttachTag(
	c *gin.Context,
) {

	var dto tagDTO.AttachTagRequest

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

	err := h.taggableService.AttachTag(
		c.Request.Context(),
		userID,
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
		http.StatusCreated,
		"tag attached successfully",
		nil,
	)
}

func (h *TaggableHandler) RemoveTag(
	c *gin.Context,
) {

	var dto tagDTO.RemoveTagRequest

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

	err := h.taggableService.RemoveTag(
		c.Request.Context(),
		userID,
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
		"tag removed successfully",
		nil,
	)
}

func (h *TaggableHandler) GetTagsByItem(
	c *gin.Context,
) {

	itemType := c.Param(
		"itemType",
	)

	itemID, err := uuid.Parse(
		c.Param("itemId"),
	)

	if err != nil {

		responses.Error(
			c,
			http.StatusBadRequest,
			"invalid item id",
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

	tags, err := h.taggableService.GetTagsByItem(
		c.Request.Context(),
		userID,
		itemType,
		itemID,
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
		"tags retrieved successfully",
		tags,
	)
}