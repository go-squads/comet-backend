package handler

import (
	"encoding/json"
	"net/http"

	"github.com/go-squads/comet-backend/repository"
	"github.com/gorilla/mux"
)

func ReadHistoryConfiguration(w http.ResponseWriter, r *http.Request) {
	historyRepo := repository.NewConfigurationRepository()
	params := mux.Vars(r)

	history := historyRepo.ReadHistory(params["app"], params["namespace"])
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(history)
}
