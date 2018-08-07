package handler

import (
	"encoding/json"
	"net/http"

	"github.com/go-squads/comet-backend/repository"
	"fmt"
)

func GetListOfApplication(w http.ResponseWriter, r *http.Request) {
	listApplication := repository.NewApplicationRepository()
	application := listApplication.GetApplicationNamespace()
	fmt.Println(listApplication)
	header :=  r.Header.Get("Authorization")
	fmt.Println(header)

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Authorization",header)
	//if header == "" {
	//	w.WriteHeader(http.StatusUnauthorized)
	//	json.NewEncoder(w).Encode(domain.Response{Status: http.StatusUnauthorized, Message: "Unauthorized"})
	//}else{
	//	w.WriteHeader(http.StatusOK)
	//}
	json.NewEncoder(w).Encode(application)
}

