package server

import (
	"fmt"
	"github.com/PlegunovN/authenticationProject/internal/server/handlers"
	"github.com/PlegunovN/authenticationProject/internal/users"
	"github.com/gorilla/mux"
	"go.uber.org/zap"
	"net/http"
)

func AuthStart(storage *users.Service, sLogger *zap.SugaredLogger) {
	api := handlers.Api{
		Storage: storage,
		SLogger: sLogger,
	}

	r := mux.NewRouter()

	fmt.Println("server start at 8081")
	sLogger.Info("hi Auth")

	r.HandleFunc("/sign_up", api.SignUp).Methods("POST")
	r.HandleFunc("/delete/{id}", api.DeleteUser).Methods("DELETE")

	err := http.ListenAndServe("127.0.0.1:8081", r)
	sLogger.Fatalf("server auth dont start, err: %w", err)
}
