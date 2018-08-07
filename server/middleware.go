package server

import (
	"net/http"
	"github.com/go-squads/comet-backend/repository"
	"github.com/go-squads/comet-backend/domain"
	"encoding/json"
)

func WithAuth(n http.HandlerFunc) http.HandlerFunc{
	return 	func(w http.ResponseWriter, r *http.Request) {

		token := r.Header.Get("Authorization")

		auth := repository.GetUserRepository().ValidateUserToken(token)
		if auth != false {
			n.ServeHTTP(w,r)
		}else{
			w.WriteHeader(http.StatusUnauthorized)
			json.NewEncoder(w).Encode(domain.Response{Status: http.StatusUnauthorized, Message: "Unauthorized"})
			return
		}
	}
}
