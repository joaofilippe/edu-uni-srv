package tokengenerator

import (
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
