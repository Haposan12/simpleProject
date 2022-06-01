package oauth

import (
	"github.com/dgrijalva/jwt-go"
	"project/freelance/danspro/simpleProject/app/models"
	"time"
)

func GenerateToken(user models.User) (string, error) {
	var mySigningKey = []byte("secretkeyjwt")
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)

	expiredAt := time.Now().AddDate(0, 0, 24)

	claims["authorized"] = true
	claims["id"] = user.Id
	claims["username"] = user.Name
	claims["exp"] = expiredAt.Unix()

	return token.SignedString(mySigningKey)
}
