package middleware

import (
	"github.com/PlegunovN/authenticationProject/internal/users"
	"github.com/golang-jwt/jwt"
	"github.com/gorilla/context"
	"go.uber.org/zap"
	"net/http"
	"strings"
)

// аутентификация
func AuthMW(next http.HandlerFunc, logger *zap.SugaredLogger) http.HandlerFunc {
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
			logger.Errorf("Error parsing token: %w", err)
			w.WriteHeader(http.StatusForbidden)
			return
		}

		claims := parsedToken.Claims.(jwt.MapClaims)

		login := claims["login"].(string)

		_, err = users.VerifyToken(token)
		if err != nil {
			logger.Errorf("Verify Token error: %w", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		//в реквест записать login из токена
		context.Set(r, "login", login)
		next(w, r)
	}
}
