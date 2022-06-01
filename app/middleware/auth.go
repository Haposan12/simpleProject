package middleware

import (
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"net/http"
	"project/freelance/danspro/simpleProject/shared/response"
	"strings"
	"time"
)

func IsAuthorized(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var (
			err error
		)
		headerAuth := r.Header.Get("Authorization")

		authorization := strings.Split(headerAuth, " ")

		if len(authorization) != 2 {
			err = errors.New("invalid authorization")
			response.SendErrorResponse(w, err.Error(), http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
			return
		} else if authorization[0] != "Bearer" {
			err = errors.New("invalid token head name")
			response.HTTPResponse(w, err.Error(), http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
			return
		}

		var mySigningKey = []byte("secretkeyjwt")
		token, err := jwt.Parse(authorization[1], func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("there was an error in parsing")
			}
			return mySigningKey, nil
		})

		if err != nil {
			err = errors.New("invalid jwt token")
			response.HTTPResponse(w, err.Error(), http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
			return
		}

		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			if claims["exp"] == nil {
				err = errors.New("missing field exp")
				response.SendErrorResponse(w, err.Error(), http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
				return
			}

			if int64(claims["exp"].(float64)) < time.Now().Unix() {
				err = errors.New("token expired")
				response.SendErrorResponse(w, err.Error(), http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
				return
			}
		}
		next.ServeHTTP(w, r)
	})
}
