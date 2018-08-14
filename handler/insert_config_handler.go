package handler

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/go-squads/comet-backend/domain"
	"github.com/go-squads/comet-backend/repository"
	"github.com/gorilla/mux"
)

func InsertConfigurationHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	decoder := json.NewDecoder(r.Body)
	header := r.Header.Get("Authorization")

	var data []domain.Configuration

	err := decoder.Decode(&data)
	if err != nil {
		log.Fatalf(err.Error())
	}

	var newConfigRequest = domain.ConfigurationRequest{
		AppName:   params["appName"],
		Namespace: params["namespaceName"],
		Data:      data,
	}

	configurationRepo := repository.NewConfigurationRepository()

	response := configurationRepo.InsertConfiguration(newConfigRequest, header)

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Authorization", header)
	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(response)
}
