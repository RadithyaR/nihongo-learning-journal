package jwt

import (
	"time"

	jwtlib "github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

func GenerateAccessToken(
	userID uuid.UUID,
	email string,
	secret string) (string, error){
		claims := CustomClaims{
			UserID: userID,
			Email: email,
			RegisteredClaims: jwtlib.RegisteredClaims{
				ExpiresAt: jwtlib.NewNumericDate(
					time.Now().Add(15 * time.Minute),
				),
			},
		}
		token := jwtlib.NewWithClaims(
			jwtlib.SigningMethodES256,
			claims,
		)

		return token.SignedString([]byte(secret))
	}