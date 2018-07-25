package server

import (
	"github.com/go-squads/comet-backend/handler"
	"github.com/gorilla/mux"
)

func Router() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/ping", handler.PingHandler).Methods("GET")
	router.HandleFunc("/configuration/{app}/{namespace}", handler.ReadConfigurationHandler).Methods("GET")
	router.HandleFunc("/configuration", handler.InsertConfigurationHandler).Methods("POST")
	router.HandleFunc("/history/{app}/{namespace}", handler.ReadHistoryConfiguration).Methods("GET")
	router.HandleFunc("/application", handler.GetListOfApplication).Methods("GET")
	return router
}
