package handler

import (
	"encoding/json"
	"github.com/go-squads/comet-backend/domain"
	"github.com/go-squads/comet-backend/repository"
	"log"
	"net/http"
)

func InsertNewApplication(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	header := r.Header.Get("Authorization")

	var newAppRequest domain.CreateApplication
	err := decoder.Decode(&newAppRequest)
	if err != nil {
		log.Fatalf(err.Error())
	}

	configurationRepo := repository.NewApplicationRepository()

	response := configurationRepo.CreateApplication(newAppRequest, header)

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Authorization", header)
	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(response)
}
