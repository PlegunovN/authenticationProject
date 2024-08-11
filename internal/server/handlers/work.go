package handlers

import (
	"encoding/json"
	"go.uber.org/zap"
	"net/http"
)

func (a Api) Work(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	err := a.validateToken(w, r, zap.SugaredLogger{})
	if err != nil {
		w.WriteHeader(http.StatusForbidden)
		return
	}

	req, err := a.storage.Work()
	if err != nil {
		return
	}
	err = json.NewEncoder(w).Encode(req)
	if err != nil {
		return
	}
	w.WriteHeader(http.StatusOK)
	return
}
