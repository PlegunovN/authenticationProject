package handlers

import (
	"encoding/json"
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

	user, token, err := a.storage.SignIn(ctx, login, password)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		a.logger.Errorf("error select books: %w", err)
		return
	}

	if user == nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	err = json.NewEncoder(w).Encode(user)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		a.logger.Errorf("error Encode books in getbooks.go: %w", err)
		return
	}
	err = json.NewEncoder(w).Encode(token)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		a.logger.Errorf("error Encode books in getbooks.go: %w", err)
		return
	}
	w.WriteHeader(http.StatusOK)
}
