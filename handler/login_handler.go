package handler

import (
	"encoding/json"
	"net/http"
	"log"
	"time"

	"github.com/go-squads/comet-backend/domain"
	"github.com/go-squads/comet-backend/repository"
)

const (
	invalidCredentialsString = "Invalid Credentials"
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
	invalidCredentialsResponse := domain.Response{Status: http.StatusUnauthorized, Message: "Invalid Credentials"}
	validCredentialsResponse := domain.Response{Status: http.StatusOK, Message: "Valid"}

	decoder := json.NewDecoder(r.Body)
	var user domain.User
	err := decoder.Decode(&user)
	if err != nil {
		log.Fatalf(err.Error())
	}

	userRepo := repository.GetUserRepository()

	token := userRepo.LogIn(user)

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
