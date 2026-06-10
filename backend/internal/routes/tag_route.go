package routes

import (
	"github.com/RadithyaR/nihongo-learning-journal/backend/internal/handlers"
	"github.com/RadithyaR/nihongo-learning-journal/backend/internal/middlewares"
	"github.com/gin-gonic/gin"
)

func TagRoute(
	router *gin.RouterGroup,
	tagHandler *handlers.TagHandler,
	taggableHandler *handlers.TaggableHandler,
) {

	tag := router.Group("/tags")

	tag.Use(middlewares.JWTMiddleware())

	tag.POST("",tagHandler.CreateTag)

	tag.GET("",tagHandler.GetTags)

	tag.PUT("/:id",tagHandler.UpdateTag)

	tag.DELETE("/:id",tagHandler.DeleteTag)

	tag.POST("/attach",taggableHandler.AttachTag)

	tag.DELETE("/attach",taggableHandler.RemoveTag)

	tag.GET("/item/:itemType/:itemId",taggableHandler.GetTagsByItem)
}