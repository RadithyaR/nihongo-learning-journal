package handlers

import (
	"log"
	"net/http"

	authDTO "github.com/RadithyaR/nihongo-learning-journal/backend/internal/dto/auth"
	serviceInterface "github.com/RadithyaR/nihongo-learning-journal/backend/internal/repositories/interfaces"
	customErrors "github.com/RadithyaR/nihongo-learning-journal/backend/pkg/errors"
	"github.com/RadithyaR/nihongo-learning-journal/backend/pkg/responses"
	appValidator "github.com/RadithyaR/nihongo-learning-journal/backend/pkg/validator"
	"github.com/gin-gonic/gin"
)

type AuthHandler struct {
	authService serviceInterface.AuthService
}

func NewAuthHandler(authService serviceInterface.AuthService) *AuthHandler{
	return &AuthHandler{
		authService: authService,
	}
}

func (h *AuthHandler) Register(c *gin.Context) {
	var dto authDTO.RegisterDTO

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

	user, err := h.authService.Register(
		c.Request.Context(),
		dto,
	)

	if err != nil{
		switch err{
		case customErrors.ErrEmailAlreadyExists:
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
		"register success",
		user,
	)
}

func (h *AuthHandler) Login(c *gin.Context){
	var dto authDTO.LoginDTO

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

	result, err := h.authService.Login(
		c.Request.Context(),
		dto,
	)

	if err != nil {
		log.Println("LOGIN ERROR:", err)
		switch err {

		case customErrors.ErrInvalidCredentials:
			responses.Error(
				c,
				http.StatusUnauthorized,
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

	c.SetCookie(
		"refresh_token",
		result.RefreshToken,
		60*60*24*7,
		"/",
		"",
		false,
		true,
	)

	responses.Success(
		c,
		http.StatusOK,
		"login success",
		gin.H{
			"access_token": result.AccessToken,
			"user": result.User,
		},
	)
}

func (h *AuthHandler) Me(c *gin.Context) {
	userID, exists := c.Get(
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

	responses.Success(
		c,
		http.StatusOK,
		"current user",
		gin.H{
			"user_id" : userID,
			"email": c.GetString("email"),
		},
	)
}

func (h *AuthHandler) RefreshToken(
	c *gin.Context,
) {

	refreshToken, err := c.Cookie(
		"refresh_token",
	)

	if err != nil {
		responses.Error(
			c,
			http.StatusUnauthorized,
			"refresh token not found",
		)

		return
	}

	result, err := h.authService.RefreshToken(
		c.Request.Context(),
		refreshToken,
	)

	if err != nil {

		switch err {

		case customErrors.ErrInvalidRefreshToken:
			responses.Error(
				c,
				http.StatusUnauthorized,
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
		"token refreshed",
		gin.H{
			"access_token": result.AccessToken,
			"user": result.User,
		},
	)
}