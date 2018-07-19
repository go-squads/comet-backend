package server

import (
	"github.com/gorilla/mux"
	"github.com/go-squads/comet-backend/handler"
)

func Router() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/ping", handler.PingHandler).Methods("GET")
	router.HandleFunc("/configuration", handler.ReadConfigurationHandler).Methods("GET")
	return router
}
