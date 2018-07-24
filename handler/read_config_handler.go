package handler

import (
	"encoding/json"
	"net/http"

	"github.com/go-squads/comet-backend/repository"
	"github.com/gorilla/mux"
)

func ReadConfigurationHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	version := r.FormValue("version")

	configurationRepo := repository.NewConfigurationRepository()

	appCfg := configurationRepo.GetConfiguration(params["app"], params["namespace"], version)

	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(appCfg)
}
