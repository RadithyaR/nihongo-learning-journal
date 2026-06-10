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

type TagHandler struct {
	tagService serviceInterfaces.TagService
}

func NewTagHandler(
	tagService serviceInterfaces.TagService,
) *TagHandler {
	return &TagHandler{
		tagService: tagService,
	}
}

func (h *TagHandler) CreateTag(
	c *gin.Context,
) {

	var dto tagDTO.CreateTagRequest

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

	tag, err := h.tagService.CreateTag(
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
		"tag created successfully",
		tag,
	)
}

func (h *TagHandler) GetTags(
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

	tags, err := h.tagService.GetTags(
		c.Request.Context(),
		userID,
	)

	if err != nil {

		responses.Error(
			c,
			http.StatusInternalServerError,
			"failed to retrieve tags",
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

func (h *TagHandler) UpdateTag(
	c *gin.Context,
) {

	tagID, err := uuid.Parse(
		c.Param("id"),
	)

	if err != nil {

		responses.Error(
			c,
			http.StatusBadRequest,
			"invalid tag id",
		)

		return
	}

	var dto tagDTO.UpdateTagRequest

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

	tag, err := h.tagService.UpdateTag(
		c.Request.Context(),
		userID,
		tagID,
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
		"tag updated successfully",
		tag,
	)
}

func (h *TagHandler) DeleteTag(
	c *gin.Context,
) {

	tagID, err := uuid.Parse(
		c.Param("id"),
	)

	if err != nil {

		responses.Error(
			c,
			http.StatusBadRequest,
			"invalid tag id",
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

	if err := h.tagService.DeleteTag(
		c.Request.Context(),
		userID,
		tagID,
	); err != nil {

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
		"tag deleted successfully",
		nil,
	)
}

