package handler

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/go-squads/comet-backend/domain"
	"github.com/go-squads/comet-backend/repository"
	"github.com/gorilla/mux"
)

func RollbackConfigurationVersion(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	decoder := json.NewDecoder(r.Body)
	header := r.Header.Get("Authorization")

	var rollbackConfig domain.ConfigurationRollback

	err := decoder.Decode(&rollbackConfig)
	if err != nil {
		log.Fatalf(err.Error())
	}

	rollbackConfig.Appname = params["appName"]
	rollbackConfig.NamespaceName = params["namespaceName"]

	fmt.Println(&rollbackConfig)
	rollback := repository.NewConfigurationRepository()
	rollbackResponse := rollback.RollbackVersionNamespace(rollbackConfig, header)

	w.Header().Set("Content-type", "application/json")
	w.Header().Set("Authorization", header)
	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(rollbackResponse)
}
