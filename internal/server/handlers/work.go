package handlers

import (
	"encoding/json"
	"net/http"
)

func (a Api) Work(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	ctx := r.Context()

	login := r.URL.Query().Get("login")
	if login == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	token := r.URL.Query().Get("token")
	if token == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	resp, err := a.Storage.Work(ctx, login, token)

	err = json.NewEncoder(w).Encode(resp)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		a.SLogger.Errorf("error select work: %w", err)
		return
	}

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		a.SLogger.Errorf("error Encode text in work.go: %w", err)
		return
	}
	w.WriteHeader(http.StatusOK)

}
