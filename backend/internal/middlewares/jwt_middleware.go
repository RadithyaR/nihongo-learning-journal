package middlewares

import (
	"net/http"
	"strings"

	"github.com/RadithyaR/nihongo-learning-journal/backend/internal/config"
	customJWT "github.com/RadithyaR/nihongo-learning-journal/backend/pkg/jwt"
	"github.com/gin-gonic/gin"
	jwtlib "github.com/golang-jwt/jwt/v5"
)

func JWTMiddleware() gin.HandlerFunc{
	return func(c *gin.Context) {
		authHeader := c.GetHeader(
			"Authorization",
		)

		if authHeader == ""{
			c.JSON(
				http.StatusUnauthorized,
				gin.H{
					"success": false,
					"message": "missing authorization header",
				},
			)
			c.Abort();
			return
		}
		tokenString := strings.TrimPrefix(
			authHeader,
			"Bearer ",
		)

		token, err := jwtlib.ParseWithClaims(
			tokenString,
			&customJWT.CustomClaims{},
			func(token *jwtlib.Token) (interface{}, error) {
				return []byte(
					config.GetEnv("JWT_SECRET"),
				), nil
			},
		)

		if err != nil{
			c.JSON(
				http.StatusUnauthorized,
				gin.H{
					"success": false,
					"message": "invalid token",
				},
			)
			c.Abort();
			return
		}

		claims, ok := token.Claims.(*customJWT.CustomClaims)

		if !ok || !token.Valid {
			c.JSON(
				http.StatusUnauthorized,
				gin.H{
					"success": false,
					"message": "invalid token",
				},
			)
			c.Abort();
			return
		}

		c.Set(
			"user_id",
			claims.UserID,
		)

		c.Set(
			"email",
			claims.Email,
		)

		c.Next()
	}
}