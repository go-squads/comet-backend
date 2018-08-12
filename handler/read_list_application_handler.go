package handler

import (
	"encoding/json"
	"net/http"

	"fmt"
	"github.com/go-squads/comet-backend/repository"
)

func GetListOfApplication(w http.ResponseWriter, r *http.Request) {
	header := r.Header.Get("Authorization")

	listApplication := repository.NewApplicationRepository()
	fmt.Println(listApplication)
	application := listApplication.GetApplicationOnly(header)
	fmt.Println(header)

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Authorization", header)
	json.NewEncoder(w).Encode(application)
}
