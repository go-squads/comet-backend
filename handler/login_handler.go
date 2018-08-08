package handler

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/go-squads/comet-backend/domain"
	"github.com/go-squads/comet-backend/repository"
	"fmt"
)

func addCookie(w http.ResponseWriter, name string, value string) {
	expire := time.Now().AddDate(0, 0, 1)
	cookie := http.Cookie{
		Name:    name,
		Value:   value,
		Expires: expire,
	}
	http.SetCookie(w, &cookie)
}

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var user domain.User
	err := decoder.Decode(&user)
	if err != nil {
		log.Fatalf(err.Error())
	}

	userRepo := repository.GetUserRepository()
	token, fullName, role := userRepo.LogIn(user)

	fmt.Println(fullName+" and "+role)

	invalidCredentialsResponse := domain.LoginResponse{Status: http.StatusUnauthorized, Fullname: "", RoleBased: "", Message: "Invalid Credentials", Token: ""}
	validCredentialsResponse := domain.LoginResponse{Status: http.StatusOK,Fullname: fullName, RoleBased: role, Message: "log_in", Token: token}

	if token == "" {
		w.WriteHeader(http.StatusUnauthorized)
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(invalidCredentialsResponse)
	} else {
		w.WriteHeader(http.StatusOK)
		addCookie(w, "token", token)
		json.NewEncoder(w).Encode(validCredentialsResponse)
	}
}
