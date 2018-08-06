package handler

import (
	"net/http"
	"encoding/json"
	"log"
	"github.com/go-squads/comet-backend/domain"
	"github.com/go-squads/comet-backend/repository"
)

func InsertNewApplication(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var newAppRequest domain.CreateApplication
	err := decoder.Decode(&newAppRequest)
	if err != nil {
		log.Fatalf(err.Error())
	}

	configurationRepo := repository.NewConfigurationRepository()

	response := configurationRepo.CreateApplication(newAppRequest)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(response)
}

