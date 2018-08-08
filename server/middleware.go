package server

import (
	"github.com/go-squads/comet-backend/domain"
	"github.com/go-squads/comet-backend/repository"
	"net/http"

	"encoding/json"
)

func WithAuth(n http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		token := r.Header.Get("Authorization")

		auth := repository.GetUserRepository().ValidateUserToken(token)
		if auth != false {
			n.ServeHTTP(w, r)
			return
		} else {
			w.WriteHeader(http.StatusUnauthorized)
			json.NewEncoder(w).Encode(domain.Response{Status: http.StatusUnauthorized, Message: "User Unauthorized"})
			return
		}
	}
}
