package handler

import (
	"encoding/json"
	"net/http"

	"github.com/go-squads/comet-backend/repository"
	"github.com/gorilla/mux"
)

func ReadConfigurationHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	//version := r.FormValue("version")
	header := r.Header.Get("Authorization")

	configurationRepo := repository.NewConfigurationRepository()

	appCfg := configurationRepo.GetConfiguration(params["appName"], params["namespaceName"], params["configVersion"], header)

	w.Header().Set("Content-type", "application/json")
	w.Header().Set("Authorization", header)
	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(appCfg)
}
