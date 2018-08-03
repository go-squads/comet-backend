package handler

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/go-squads/comet-backend/domain"
	"github.com/go-squads/comet-backend/repository"
)

func RollbackConfigurationVersion(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var rollbackConfig domain.ConfigurationRollback

	err := decoder.Decode(&rollbackConfig)
	if err != nil {
		log.Fatalf(err.Error())
	}

	fmt.Println(&rollbackConfig)
	rollback := repository.NewConfigurationRepository()
	rollbackResponse := rollback.RollbackVersionNamespace(rollbackConfig)

	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(rollbackResponse)
}
