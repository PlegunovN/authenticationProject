package server

import (
	"fmt"
	"github.com/PlegunovN/authenticationProject/internal/server/handlers"
	"github.com/PlegunovN/authenticationProject/internal/users"
	"github.com/gorilla/mux"
	"go.uber.org/zap"
	"net/http"
	"time"
)

func AuthStart(storage *users.Service, logger *zap.SugaredLogger) {

	api := handlers.New(storage, logger)

	r := mux.NewRouter()

	handler := loggingMiddleware(r, logger)

	logger.Info("server start at 8081")

	r.HandleFunc("/sign_up", api.SignUp).Methods("POST")
	r.HandleFunc("/delete/{id}", api.DeleteUser).Methods("DELETE")
	r.HandleFunc("/sign_in", api.SignIn).Methods("GET")
	r.HandleFunc("/work", api.Work).Methods("GET")

	err := http.ListenAndServe("127.0.0.1:8081", handler)
	logger.Fatalf("server auth dont start, err: %w", err, handler)

}
func loggingMiddleware(next http.Handler, logger *zap.SugaredLogger) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		next.ServeHTTP(w, r)
		message := fmt.Sprintf("%s %s %s", r.Method, r.RequestURI, time.Since(start))
		logger.Info(message)
	})
}
