package handlers

import (
	"net/http"
)

func (a Api) DeleteUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	ctx := r.Context()

	login := r.URL.Query().Get("login")
	if login == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err := a.storage.DeleteUser(ctx, login)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		a.logger.Errorf("error Encode id in delete.go: %w", err)
		return
	}

	w.WriteHeader(http.StatusOK)
}
