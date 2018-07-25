package handler

import (
	"encoding/json"
	"net/http"

	"github.com/go-squads/comet-backend/repository"
)

func GetListOfApplication(w http.ResponseWriter, r *http.Request) {
	listApplication := repository.NewConfigurationRepository()
	application := listApplication.GetApplicationNamespace()

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(application)
}
