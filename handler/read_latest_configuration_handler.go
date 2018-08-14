package handler

import (
	"encoding/json"
	"net/http"

	"github.com/go-squads/comet-backend/repository"
	"github.com/gorilla/mux"
)

func GetLatestConfiguration(w http.ResponseWriter, r *http.Request) {
	header := r.Header.Get("Authorization")
	params := mux.Vars(r)

	configurationRepo := repository.NewConfigurationRepository()

	response := configurationRepo.GetLatestConfiguration(params["appName"], params["namespaceName"], header)

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Authorization", header)
	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(response)
}
