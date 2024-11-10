package tokenmanager

import (
	"fmt"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func GenerateToken(claims ICustomClaim) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userID":     claims.ID(),
		"userTypeID": claims.UserDetails().ID(),
		"userType":   claims.UserType().String(),
		"exp":        time.Now().Add(time.Hour * 72).Unix(),
	})

	return token.SignedString([]byte(os.Getenv("SECRET_KEY")))
}

func ValidateToken(tokenString string) (jwt.MapClaims, error) {
	secretKey := []byte(os.Getenv("SECRET_KEY"))
	claims := jwt.MapClaims{}
	raw, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return secretKey, nil
	})

	if err != nil {
		return nil, err
	}

	if !raw.Valid {
		return nil, fmt.Errorf("invalid token")
	}

	return claims, nil
}
