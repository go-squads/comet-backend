package server

import (
	"net/http"
	"log"
)
func StartAPIServer(){
	muxRouter := Router()
	log.Fatal(http.ListenAndServe(":8000",muxRouter))
}
