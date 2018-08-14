package handler

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"fmt"

	"github.com/go-squads/comet-backend/domain"
	"github.com/go-squads/comet-backend/repository"
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

	if len(fullName) > 0 {
		fmt.Println(fullName + " and " + role)

		validCredentialsResponse := domain.LoginResponse{Status: http.StatusOK, Fullname: fullName, RoleBased: role, Message: "log_in", Token: token}

		w.WriteHeader(http.StatusOK)
		addCookie(w, "token", token)
		json.NewEncoder(w).Encode(validCredentialsResponse)
	} else {
		fmt.Println("not found")

		invalidCredentialsResponse := domain.LoginResponse{Status: http.StatusUnauthorized, Fullname: "", RoleBased: "", Message: "Invalid Credentials", Token: ""}

		w.WriteHeader(http.StatusOK)
		addCookie(w, "token", token)
		json.NewEncoder(w).Encode(invalidCredentialsResponse)
	}

}
