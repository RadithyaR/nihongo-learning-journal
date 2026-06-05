package routes

import (
	"github.com/RadithyaR/nihongo-learning-journal/backend/internal/handlers"
	"github.com/RadithyaR/nihongo-learning-journal/backend/internal/middlewares"
	"github.com/gin-gonic/gin"
)

func AuthRoute(router *gin.RouterGroup, authHandler *handlers.AuthHandler) {
	auth := router.Group("/auth")

	auth.POST("/register", authHandler.Register)
	auth.POST("/verify-email", authHandler.VerifyEmail)
	auth.POST("/login", authHandler.Login)

	auth.POST("/forgot-password", authHandler.ForgotPassword)
	auth.POST("/reset-password", authHandler.ResetPassword)
	
	auth.POST("/refresh", authHandler.RefreshToken)
	auth.POST("/logout", authHandler.Logout)
	
	auth.POST("/change-password", middlewares.JWTMiddleware(), authHandler.ChangePassword)
	auth.POST("/logout-all", middlewares.JWTMiddleware(), authHandler.LogoutAll)

	auth.GET("/me", middlewares.JWTMiddleware(), authHandler.Me)
}