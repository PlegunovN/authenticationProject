package middleware

import (
	"fmt"
	"go.uber.org/zap"
	"net/http"
	"time"
)

func LoggingMiddleware(next http.Handler, logger *zap.SugaredLogger) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		next.ServeHTTP(w, r)
		message := fmt.Sprintf("%s %s %s", r.Method, r.RequestURI, time.Since(start))
		logger.Info(message)
	})
}
