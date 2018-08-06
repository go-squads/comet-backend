package handler

import (
	"encoding/json"
	"github.com/go-squads/comet-backend/domain"
	"github.com/go-squads/comet-backend/repository"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"fmt"
)

func InsertNewNamespace(w http.ResponseWriter, r *http.Request) {
	var newNamespace domain.Namespace

	params := mux.Vars(r)
	decoder := json.NewDecoder(r.Body)

	err := decoder.Decode(&newNamespace)
	if err != nil {
		log.Fatalf(err.Error())
	}

	applicationRepo := repository.NewApplicationRepository()

	fmt.Println(params["appName"])
	fmt.Println(newNamespace)
	newNamespaces := applicationRepo.CreateNewNamespace(params["appName"], newNamespace)

	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(newNamespaces)
}
