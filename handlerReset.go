package main

import (
	"errors"
	"net/http"
)

func (cfg *apiConfig) handlerResetMetrics(w http.ResponseWriter, r *http.Request) {

	if cfg.platform != "dev" {
		w.WriteHeader(http.StatusForbidden)
		respondWithError(w, 403, "forbidden", errors.New("forbidden"))
		return
	}

	cfg.fileserverHits.Store(0)
	err := cfg.db.DeleteAllUsers(r.Context())
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		respondWithError(w, 500, "error deleting users", err)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Hits reset to 0 and database reset."))
}
