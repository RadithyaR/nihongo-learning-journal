package responses

import "github.com/gin-gonic/gin"

func Error(
	c *gin.Context,
	status int,
	message string,
) {
	c.JSON(status, gin.H{
		"success": false,
		"message": message,
	})
	
}