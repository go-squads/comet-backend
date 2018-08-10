package server

import (
	"github.com/go-squads/comet-backend/handler"
	"github.com/gorilla/mux"
)

func Router() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/ping", handler.PingHandler).Methods("GET")
	router.HandleFunc("/configuration", WithAuth(handler.InsertConfigurationHandler)).Methods("POST")
	router.HandleFunc("/configuration/rollback", WithAuth(handler.RollbackConfigurationVersion)).Methods("POST")
	router.HandleFunc("/configuration/{app}/{namespace}", WithAuth(handler.ReadConfigurationHandler)).Methods("GET")
	router.HandleFunc("/configuration/history/{app}/{namespace}", WithAuth(handler.ReadHistoryConfiguration)).Methods("GET")
	router.HandleFunc("/application", WithAuth(handler.GetListOfApplication)).Methods("GET")
	router.HandleFunc("/application/create", WithAuth(handler.InsertNewApplication)).Methods("POST")
	router.HandleFunc("/application/create/{appName}", handler.InsertNewNamespace).Methods("POST")
	router.HandleFunc("/application/{app}/namespaces/{namespaces}/configurations",WithAuth(handler.GetLatestConfiguration)).Methods("GET")
	router.HandleFunc("/login", handler.LoginHandler).Methods("POST")
	return router
}
