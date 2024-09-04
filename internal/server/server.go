package server

import (
	"github.com/PlegunovN/authenticationProject/internal/middleware"
	"github.com/PlegunovN/authenticationProject/internal/server/handlers"
	"github.com/PlegunovN/authenticationProject/internal/users"
	"github.com/gorilla/mux"
	"go.uber.org/zap"
	"net/http"
)

func Run(service *users.Service, logger *zap.SugaredLogger) {

	api := handlers.New(service, logger)

	r := mux.NewRouter()

	handler := middleware.LoggingMiddleware(r, logger)

	logger.Info("server start at 8081")

	r.HandleFunc("/sign_up", api.SignUp).Methods("POST")
	r.HandleFunc("/delete", middleware.AuthMW(api.DeleteUser)).Methods("DELETE")
	r.HandleFunc("/sign_in", api.SignIn).Methods("GET")

	err := http.ListenAndServe("127.0.0.1:8081", handler)
	logger.Fatalf("server auth dont start, err: %w", err, handler)

}
