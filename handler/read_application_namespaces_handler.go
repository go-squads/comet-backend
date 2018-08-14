package handler

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-squads/comet-backend/repository"
	"github.com/gorilla/mux"
)

func GetListOfNamespaces(w http.ResponseWriter, r *http.Request) {
	header := r.Header.Get("Authorization")
	params := mux.Vars(r)

	listApplication := repository.NewApplicationRepository()
	fmt.Println(listApplication)
	application := listApplication.GetApplicationNamespace(header, params["appName"])
	fmt.Println(header)

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Authorization", header)
	json.NewEncoder(w).Encode(application)
}
