package handlers

import (
	"net/http"

	goalDTO "github.com/RadithyaR/nihongo-learning-journal/backend/internal/dto/goal"
	serviceInterfaces "github.com/RadithyaR/nihongo-learning-journal/backend/internal/services/interfaces"
	customErrors "github.com/RadithyaR/nihongo-learning-journal/backend/pkg/errors"
	"github.com/RadithyaR/nihongo-learning-journal/backend/pkg/responses"
	appValidator "github.com/RadithyaR/nihongo-learning-journal/backend/pkg/validator"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type GoalHandler struct {
	goalService serviceInterfaces.GoalService
}

func NewGoalHandler(
	goalService serviceInterfaces.GoalService,
) *GoalHandler {
	return &GoalHandler{
		goalService: goalService,
	}
}

func (h *GoalHandler) CreateGoal(
	c *gin.Context,
) {

	var dto goalDTO.CreateGoalRequest

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

	goal, err :=
		h.goalService.CreateGoal(
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
		"goal created successfully",
		goal,
	)
}

func (h *GoalHandler) GetGoalList(
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

	goals, err :=
		h.goalService.GetGoalList(
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
		"goals retrieved successfully",
		goals,
	)
}

func (h *GoalHandler) GetGoalByID(
	c *gin.Context,
) {

	idParam := c.Param("id")

	goalID, err := uuid.Parse(
		idParam,
	)

	if err != nil {

		responses.Error(
			c,
			http.StatusBadRequest,
			"invalid goal id",
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

	goal, err :=
		h.goalService.GetGoalByID(
			c.Request.Context(),
			userID,
			goalID,
		)

	if err != nil {

		switch err {

		case customErrors.ErrGoalNotFound:

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
		"goal retrieved successfully",
		goal,
	)
}

func (h *GoalHandler) UpdateGoal(
	c *gin.Context,
) {

	idParam := c.Param("id")

	goalID, err := uuid.Parse(
		idParam,
	)

	if err != nil {

		responses.Error(
			c,
			http.StatusBadRequest,
			"invalid goal id",
		)

		return
	}

	var dto goalDTO.UpdateGoalRequest

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

	goal, err :=
		h.goalService.UpdateGoal(
			c.Request.Context(),
			userID,
			goalID,
			dto,
		)

	if err != nil {

		switch err {

		case customErrors.ErrGoalNotFound:

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
		"goal updated successfully",
		goal,
	)
}

func (h *GoalHandler) CompleteGoal(
	c *gin.Context,
) {

	idParam := c.Param("id")

	goalID, err := uuid.Parse(
		idParam,
	)

	if err != nil {

		responses.Error(
			c,
			http.StatusBadRequest,
			"invalid goal id",
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

	err = h.goalService.CompleteGoal(
		c.Request.Context(),
		userID,
		goalID,
	)

	if err != nil {

		switch err {

		case customErrors.ErrGoalNotFound:

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
		"goal completed successfully",
		nil,
	)
}

func (h *GoalHandler) CancelGoal(
	c *gin.Context,
) {

	idParam := c.Param("id")

	goalID, err := uuid.Parse(
		idParam,
	)

	if err != nil {

		responses.Error(
			c,
			http.StatusBadRequest,
			"invalid goal id",
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

	err = h.goalService.CancelGoal(
		c.Request.Context(),
		userID,
		goalID,
	)

	if err != nil {

		switch err {

		case customErrors.ErrGoalNotFound:

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
		"goal cancelled successfully",
		nil,
	)
}

func (h *GoalHandler) DeleteGoal(
	c *gin.Context,
) {

	idParam := c.Param("id")

	goalID, err := uuid.Parse(
		idParam,
	)

	if err != nil {

		responses.Error(
			c,
			http.StatusBadRequest,
			"invalid goal id",
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

	err = h.goalService.DeleteGoal(
		c.Request.Context(),
		userID,
		goalID,
	)

	if err != nil {

		switch err {

		case customErrors.ErrGoalNotFound:

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
		"goal deleted successfully",
		nil,
	)
}