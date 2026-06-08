package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func GetUserID(
	c *gin.Context,
) (uuid.UUID, bool) {

	userIDValue, exists := c.Get(
		"user_id",
	)

	if !exists {
		return uuid.Nil, false
	}

	userID, ok := userIDValue.(uuid.UUID)

	if !ok {
		return uuid.Nil, false
	}

	return userID, true
}