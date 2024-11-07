package token

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	userentities "github.com/joaofilippe/edu-uni-srv/core/entities/user"
)

func GenerateToken(user userentities.User) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userID":     user.ID(),
		"userTypeID": user.UserDetails().ID(),
		"userType":   user.UserType(),
		"exp":        time.Now().Add(time.Hour * 72).Unix(),
	})

	return token.SignedString([]byte(os.Getenv("SECRET_KEY")))
}
