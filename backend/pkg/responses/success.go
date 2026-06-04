package responses

import "github.com/gin-gonic/gin"

func Success(
	c *gin.Context,
	status int,
	message string,
	data any,
) {
	c.JSON(status, gin.H{
		"success": true,
		"message": message,
		"data": data,
	})
}