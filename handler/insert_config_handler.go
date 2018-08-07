package handler

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/go-squads/comet-backend/domain"
	"github.com/go-squads/comet-backend/repository"
)

func InsertConfigurationHandler(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	header :=  r.Header.Get("Authorization")

	var newConfigRequest domain.ConfigurationRequest

	err := decoder.Decode(&newConfigRequest)
	if err != nil {
		log.Fatalf(err.Error())
	}

	configurationRepo := repository.NewConfigurationRepository()

	response := configurationRepo.InsertConfiguration(newConfigRequest)

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Authorization",header)
	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(response)
}
