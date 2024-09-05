package server

import (
	"github.com/PlegunovN/authenticationProject/configs"
	"github.com/PlegunovN/authenticationProject/internal/server/handlers"
	"github.com/PlegunovN/authenticationProject/internal/server/middleware"
	"github.com/PlegunovN/authenticationProject/internal/users"
	"github.com/gorilla/mux"
	"go.uber.org/zap"
	"net/http"
)

func Run(userService *users.Service, logger *zap.SugaredLogger, secretKey *configs.Config) {

	api := handlers.New(userService, logger, secretKey)

	r := mux.NewRouter()

	handler := middleware.LoggingMiddleware(r, logger)

	logger.Info("server start at 8081")

	r.HandleFunc("/sign_up", api.SignUp).Methods("POST")
	r.HandleFunc("/delete", middleware.AuthMW(api.DeleteUser, logger, secretKey)).Methods("DELETE")
	r.HandleFunc("/sign_in", api.SignIn).Methods("GET")

	err := http.ListenAndServe("127.0.0.1:8081", handler)
	if err != nil {
		logger.Fatal("server auth dont start, err: %w", err)
	}
}
