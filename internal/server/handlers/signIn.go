package handlers

import (
	"encoding/json"
	"github.com/PlegunovN/authenticationProject/internal/users"
	"net/http"
)

func (a Api) SignIn(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	ctx := r.Context()

	login := r.URL.Query().Get("login")
	if login == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	password := r.URL.Query().Get("password")
	if password == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	token, err := a.userService.SignIn(ctx, login, password)
	if err != nil {
		if _, ok := err.(users.ErrorPasswordIncorrect); ok {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		w.WriteHeader(http.StatusInternalServerError)
		a.logger.Errorf("error select books: %w", err)
		return
	}

	if token == "" {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	if err = json.NewEncoder(w).Encode(token); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		a.logger.Errorf("error Encode token: %w", err)
		return
	}

	w.WriteHeader(http.StatusOK)
}
