package server

import (
	"github.com/go-squads/comet-backend/handler"
	"github.com/gorilla/mux"
)

func Router() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/ping", handler.PingHandler).Methods("GET")

	router.HandleFunc("/login", handler.LoginHandler).Methods("POST")

	router.HandleFunc("/applications", WithAuth(handler.GetListOfApplication)).Methods("GET")
	router.HandleFunc("/applications/{appName}/namespaces", WithAuth(handler.GetListOfNamespaces)).Methods("GET")
	router.HandleFunc("/applications/{appName}/namespaces/{namespaceName}/configurations/latest", WithAuth(handler.GetLatestConfiguration)).Methods("GET")
	router.HandleFunc("/applications/{appName}/namespaces/{namespaceName}/configurations/{configVersion}", WithAuth(handler.ReadConfigurationHandler)).Methods("GET")
	router.HandleFunc("/applications/{appName}/namespaces/{namespaceName}/histories", WithAuth(handler.ReadHistoryConfiguration)).Methods("GET")

	router.HandleFunc("/applications", WithAuth(handler.InsertNewApplication)).Methods("POST")
	router.HandleFunc("/applications/{appName}/namespaces", handler.InsertNewNamespace).Methods("POST")
	router.HandleFunc("/applications/{appName}/namespaces/{namespaceName}/configurations", WithAuth(handler.InsertConfigurationHandler)).Methods("POST")
	router.HandleFunc("/applications/{appName}/namespaces/{namespaceName}/configurations/rollback", WithAuth(handler.RollbackConfigurationVersion)).Methods("POST")

	return router
}
