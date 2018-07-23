package handler

import (
	"encoding/json"
	"net/http"

	"github.com/go-squads/comet-backend/repository"
)

func ReadHistoryConfiguration(w http.ResponseWriter, r *http.Request) {
	historyRepo := repository.NewConfigurationRepository()

	history := historyRepo.ReadHistory()
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(history)
}
