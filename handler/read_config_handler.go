package handler

import (
	"encoding/json"
	"net/http"

	"github.com/go-squads/comet-backend/repository"
)

func ReadConfigurationHandler(w http.ResponseWriter, r *http.Request) {
	app := r.FormValue("app")
	namespace := r.FormValue("namespace")
	version := r.FormValue("version")

	configurationRepo := repository.NewConfigurationRepository()

	cfg := configurationRepo.GetConfiguration(app, namespace, version)

	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(cfg)
}
