package middleware

import (
	"github.com/golang-jwt/jwt"
	"github.com/gorilla/context"
	"go.uber.org/zap"
	"net/http"
	"strings"
)

// аутентификация
func AuthMW(next http.HandlerFunc, logger *zap.SugaredLogger, tokenSecretKey string) http.HandlerFunc {
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
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {

				return nil, nil
			}
			return []byte(tokenSecretKey), nil
		})
		if err != nil {
			logger.Errorf("Error parsing token: %w", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		claims := parsedToken.Claims.(jwt.MapClaims)

		login := claims["login"].(string)

		//в реквест записать login из токена
		context.Set(r, "login", login)
		next(w, r)
	}
}
