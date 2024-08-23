package middleware

import (
	"fmt"
	"github.com/PlegunovN/authenticationProject/internal/users"
	"github.com/golang-jwt/jwt"
	"github.com/gorilla/context"
	"log"
	"net/http"
	"strings"
)

// аутонтификация
func AuthMW(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		token := r.Header.Get("Authorization")
		if len(token) == 0 {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		// Проверяем формат токена
		parts := strings.SplitN(token, " ", 2)
		if parts[0] != "Bearer" {
			return
		}

		token = parts[1]

		parsedToken, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
			return users.TokenSecretKey, nil
		})
		if err != nil {
			fmt.Println("Error parsing token:", err)
			w.WriteHeader(http.StatusForbidden)
			return
		}

		claims := parsedToken.Claims.(jwt.MapClaims)

		login := claims["login"].(string)
		exp := claims["ExpiresAt"].(string)
		fmt.Println(exp)
		fmt.Println(login)
		fmt.Println(parsedToken.Claims)

		_, err = users.VerifyToken(token)
		log.Print(err)
		if err != nil {

			return
		}

		//в реквест записать login из токена
		context.Set(r, "login", login)
		next(w, r)
	}
}
