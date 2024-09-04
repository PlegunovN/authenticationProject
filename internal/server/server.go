package server

import (
	"github.com/PlegunovN/authenticationProject/internal/server/handlers"
	middleware2 "github.com/PlegunovN/authenticationProject/internal/server/middleware"
	"github.com/PlegunovN/authenticationProject/internal/users"
	"github.com/gorilla/mux"
	"go.uber.org/zap"
	"net/http"
)

func Run(userService *users.Service, logger *zap.SugaredLogger) {

	api := handlers.New(userService, logger)

	r := mux.NewRouter()

	handler := middleware2.LoggingMiddleware(r, logger)

	logger.Info("server start at 8081")

	r.HandleFunc("/sign_up", api.SignUp).Methods("POST")
	r.HandleFunc("/delete", middleware2.AuthMW(api.DeleteUser, logger)).Methods("DELETE")
	r.HandleFunc("/sign_in", api.SignIn).Methods("GET")

	err := http.ListenAndServe("127.0.0.1:8081", handler)
	if err != nil {
		logger.Fatalf("server auth dont start, err: %w", err, handler)
	}
}
