package handlers

import (
	"encoding/json"
	"net/http"
)

type CreateRequest struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}

func (a Api) SignUp(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	req := new(CreateRequest)
	ctx := r.Context()

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		a.logger.Errorf("error in decoder, create.go: %w", err)
		return
	}

	if req.Login == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if req.Password == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err = a.storage.SignUp(ctx, req.Login, req.Password)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		a.logger.Errorf("err in create.go: %w", err)
		return
	}
	w.WriteHeader(http.StatusCreated)
}
